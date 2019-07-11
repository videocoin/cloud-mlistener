package service

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	vc "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	notificationv1 "github.com/videocoin/cloud-api/notifications/v1"
	"github.com/videocoin/cloud-pkg/streamManager"
)

func interfaceByName(name string) func() interface{} {
	switch name {
	// stream manager events
	case "StreamRequested":
		return func() interface{} { return new(streamManager.ManagerStreamRequested) }
	case "StreamApproved":
		return func() interface{} { return new(streamManager.ManagerStreamApproved) }
	case "StreamCreated":
		return func() interface{} { return new(streamManager.ManagerStreamCreated) }
	case "StreamEnded":
		return func() interface{} { return new(streamManager.ManagerStreamEnded) }
	case "InputChunkAdded":
		return func() interface{} { return new(streamManager.ManagerInputChunkAdded) }
	case "RefundAllowed":
		return func() interface{} { return new(streamManager.ManagerRefundAllowed) }
	case "RefundRevoked":
		return func() interface{} { return new(streamManager.ManagerRefundRevoked) }
	case "ValidatorAdded":
		return func() interface{} { return new(streamManager.ManagerValidatorAdded) }
	case "ValidatorRemoved":
		return func() interface{} { return new(streamManager.ManagerValidatorRemoved) }
	case "OwnershipTransferred":
		return func() interface{} { return new(streamManager.ManagerOwnershipTransferred) }
	}

	return nil
}

type ListenerConfig struct {
	NodeHTTPAddr string
	RedisAddr    string
	ContractAddr string
	BlockNumber  *big.Int
	EB           *EventBus
	DS           *Datastore
	Logger       *logrus.Entry
}

type listener struct {
	ec *ethclient.Client
	rc *redis.Client
	eb *EventBus
	ds *Datastore

	contractAddr common.Address
	contract     *bind.BoundContract
	signatures   map[string]string

	blockNumber *big.Int

	logger *logrus.Entry
}

func NewListener(c *ListenerConfig) (*listener, error) {
	ec, err := ethclient.Dial(c.NodeHTTPAddr)
	if err != nil {
		return nil, err
	}

	opts, err := redis.ParseURL(c.RedisAddr)
	if err != nil {
		return nil, err
	}

	rc := redis.NewClient(opts)

	contractAddr := common.HexToAddress(c.ContractAddr)
	abiName := string(streamManager.ManagerABI)

	contractAbi, err := abi.JSON(strings.NewReader(abiName))
	if err != nil {
		return nil, err
	}

	signatures := make(map[string]string)
	for name, ev := range contractAbi.Events {
		signatures[ev.Id().Hex()] = name
	}

	contract := bind.NewBoundContract(
		contractAddr, contractAbi, ec, ec, ec)

	block, err := ec.BlockByNumber(context.Background(), c.BlockNumber)
	if err != nil {
		return nil, err
	}

	return &listener{
		ec:           ec,
		rc:           rc,
		eb:           c.EB,
		ds:           c.DS,
		logger:       c.Logger.WithField("component", "listener"),
		contract:     contract,
		contractAddr: contractAddr,
		signatures:   signatures,
		blockNumber:  block.Number(),
	}, nil
}

func (l *listener) Start() {
	var lastSeenLog types.Log
	starting := true

	for {
		query := vc.FilterQuery{
			Addresses: []common.Address{
				l.contractAddr,
			},
			FromBlock: l.blockNumber,
		}

		logs, err := l.ec.FilterLogs(context.Background(), query)
		if err != nil {
			l.logger.Error(err)
		}

		if len(logs) == 0 {
			time.Sleep(5 * time.Second)
			continue
		}

		if starting == true {
			lastSeenLog = logs[0]
		}

		for _, vLog := range logs {
			l.logger.Infof("last seen log with block number %d and index %d, tx %d ",
				lastSeenLog.BlockNumber, lastSeenLog.Index, lastSeenLog.TxIndex)

			l.logger.Infof("parsing log with block number %d and index %d, tx %d ",
				vLog.BlockNumber, vLog.Index, vLog.TxIndex)

			if starting == false && vLog.BlockNumber == lastSeenLog.BlockNumber && vLog.Index <= lastSeenLog.Index {
				l.logger.Info("skipping ...")

				time.Sleep(5 * time.Second)
				continue
			}

			starting = false

			lastSeenLog = vLog

			l.logger.Infof("updated last seen log with block number %d and index %d, tx %d ",
				lastSeenLog.BlockNumber, lastSeenLog.Index, lastSeenLog.TxIndex)

			l.blockNumber.SetUint64(vLog.BlockNumber)

			for _, topic := range vLog.Topics {
				l.logger.Infof("parsing topic %s", topic.Hex())
				name := l.signatures[topic.Hex()]
				f := interfaceByName(name)
				if f == nil {
					continue
				}

				out := f()
				if err := l.contract.UnpackLog(out, name, vLog); err != nil {
					l.logger.Error(err)
					continue
				}

				switch v := out.(type) {
				case *streamManager.ManagerStreamRequested:
					pipelineId, userId, err := l.getSetCache(v.StreamId.Uint64())
					if err != nil {
						l.logger.Warnf("StreamRequested set cache failed: %s", err.Error())
						break
					}

					l.logger.Infof("%s, %d, %s, %s, %s", "StreamRequested", v.StreamId.Uint64(), v.Client.Hex(), pipelineId, userId)

					err = l.sendNotificationEvent(
						map[string]string{
							"event":       "pipeline/stream",
							"pipeline_id": pipelineId,
							"user_id":     userId,
							"message":     fmt.Sprintf("stream id %s has been requested", v.StreamId.String()),
						},
					)
					if err != nil {
						l.logger.Warnf("sendNotificationEvent failed: %s", err.Error())
					}

				case *streamManager.ManagerStreamApproved:
					pipelineId, userId, err := l.getSetCache(v.StreamId.Uint64())
					if err != nil {
						l.logger.Warnf("StreamApproved get cache failed %s", err.Error())
						break
					}

					l.logger.Infof("%s, %d, %s, %s", "StreamApproved", v.StreamId.Uint64(), pipelineId, userId)

					err = l.sendNotificationEvent(
						map[string]string{
							"event":       "pipeline/stream",
							"pipeline_id": pipelineId,
							"user_id":     userId,
							"message":     fmt.Sprintf("stream id %s has been approved", v.StreamId.String()),
						},
					)
					if err != nil {
						l.logger.Warnf("sendNotificationEvent failed: %s", err.Error())
					}

				case *streamManager.ManagerStreamCreated:
					pipelineId, userId, err := l.getSetCache(v.StreamId.Uint64())
					if err != nil {
						l.logger.Warnf("StreamCreated get cache failed %s", err.Error())
						break
					}

					l.logger.Infof("%s, %d, %s, %s, %s",
						"StreamCreated", v.StreamId.Uint64(), v.StreamAddress.Hex(), pipelineId, userId)

					err = l.sendNotificationEvent(
						map[string]string{
							"event":       "pipeline/stream",
							"pipeline_id": pipelineId,
							"user_id":     userId,
							"message":     fmt.Sprintf("stream id %s has been created with stream address %s", v.StreamId.String(), v.StreamAddress.Hex()),
						},
					)
					if err != nil {
						l.logger.Warnf("sendNotificationEvent failed: %s", err.Error())
					}

				case *streamManager.ManagerStreamEnded:
					pipelineId, userId, err := l.getSetCache(v.StreamId.Uint64())
					if err != nil {
						l.logger.Warnf("StreamEnded get cache failed %s", err.Error())
						break
					}

					l.logger.Infof("%s, %d, %s, %s", "StreamEnded", v.StreamId.Uint64(), pipelineId, userId)

					err = l.sendNotificationEvent(
						map[string]string{
							"event":       "pipeline/stream",
							"pipeline_id": pipelineId,
							"user_id":     userId,
							"message":     fmt.Sprintf("stream id %s has been ended", v.StreamId.String()),
						},
					)
					if err != nil {
						l.logger.Warnf("sendNotificationEvent failed: %s", err.Error())
					}

				case *streamManager.ManagerInputChunkAdded:
					pipelineId, userId, err := l.getSetCache(v.StreamId.Uint64())
					if err != nil {
						l.logger.Warnf("InputChunkAdded get cache failed %s", err.Error())
						break
					}

					l.logger.Infof("%s, %d, %s, %s", "InputChunkAdded", v.StreamId.Uint64(), pipelineId, userId)

					err = l.sendNotificationEvent(
						map[string]string{
							"event":       "pipeline/stream",
							"pipeline_id": pipelineId,
							"user_id":     userId,
							"message":     fmt.Sprintf("input chunk has been added on stream id %s", v.StreamId.String()),
						},
					)
					if err != nil {
						l.logger.Warnf("sendNotificationEvent failed: %s", err.Error())
					}

				case *streamManager.ManagerRefundAllowed:
					pipelineId, userId, err := l.getSetCache(v.StreamId.Uint64())
					if err != nil {
						l.logger.Warnf("RefundAllowed get cache failed %s", err.Error())
						break
					}

					l.logger.Infof("%s, %d, %s, %s", "RefundAllowed", v.StreamId.Uint64(), pipelineId, userId)

					err = l.sendNotificationEvent(
						map[string]string{
							"event":       "pipeline/stream",
							"pipeline_id": pipelineId,
							"user_id":     userId,
							"message":     fmt.Sprintf("refund has been allowed on stream id %s", v.StreamId.String()),
						},
					)
					if err != nil {
						l.logger.Warnf("sendNotificationEvent failed: %s", err.Error())
					}

				case *streamManager.ManagerRefundRevoked:
					pipelineId, userId, err := l.getSetCache(v.StreamId.Uint64())
					if err != nil {
						l.logger.Warnf("RefundRevoked get cache failed %s", err.Error())
						break
					}

					l.logger.Infof("%s, %d, %s, %s", "RefundRevoked", v.StreamId.Uint64(), pipelineId, userId)

					err = l.sendNotificationEvent(
						map[string]string{
							"event":       "pipeline/stream",
							"pipeline_id": pipelineId,
							"user_id":     userId,
							"message":     fmt.Sprintf("refund has been revoked on stream id %s", v.StreamId.String()),
						},
					)
					if err != nil {
						l.logger.Warnf("sendNotificationEvent failed: %s", err.Error())
					}

				case *streamManager.ManagerValidatorAdded:
					l.logger.Infof("%s", "ValidatorAdded")
				case *streamManager.ManagerValidatorRemoved:
					l.logger.Infof("%s", "ValidatorRemoved")
				case *streamManager.ManagerOwnershipTransferred:
					l.logger.Infof("%s", "OwnershipTransferred")
				}
			}
		}
	}
}

func (l *listener) Stop() {
}

func (l *listener) sendNotificationEvent(params map[string]string) error {
	notification := &notificationv1.Notification{
		Target: notificationv1.NotificationTarget_WEB,
		Params: params,
	}

	err := l.eb.SendNotification(notification)
	if err != nil {
		return err
	}

	return nil
}

func (l *listener) setCache(streamID uint64) (string, string, error) {
	j, err := l.ds.GetJobByStreamId(streamID)
	if err != nil {
		return "", "", err
	}

	p, err := l.ds.GetPipelineById(j.PipelineId)
	if err != nil {
		return "", "", err
	}

	err = l.rc.SetNX(fmt.Sprintf("u-%d", streamID), p.UserId, 24*time.Hour).Err()
	if err != nil {
		return "", "", err
	}

	err = l.rc.SetNX(fmt.Sprintf("p-%d", streamID), p.Id, 24*time.Hour).Err()
	if err != nil {
		return "", "", err
	}

	return p.Id, p.UserId, nil
}

func (l *listener) getCache(streamID uint64) (string, string, error) {
	userId, err := l.rc.Get(fmt.Sprintf("u-%d", streamID)).Result()
	if err != nil {
		return "", "", err
	}

	pipelineId, err := l.rc.Get(fmt.Sprintf("p-%d", streamID)).Result()
	if err != nil {
		return "", "", err
	}

	return pipelineId, userId, nil
}

func (l *listener) getSetCache(streamID uint64) (string, string, error) {
	pId, uId, err := l.getCache(streamID)
	if err != nil {
		pId, uId, err = l.setCache(streamID)
		if err != nil {
			return "", "", err
		}
	}

	return pId, uId, nil
}
