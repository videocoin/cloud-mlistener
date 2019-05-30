package service

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	notificationv1 "github.com/VideoCoin/cloud-api/notifications/v1"
	"github.com/VideoCoin/cloud-pkg/streamManager"
	vc "github.com/VideoCoin/go-videocoin"
	"github.com/VideoCoin/go-videocoin/accounts/abi"
	"github.com/VideoCoin/go-videocoin/accounts/abi/bind"
	"github.com/VideoCoin/go-videocoin/common"
	"github.com/VideoCoin/go-videocoin/ethclient"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
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

	rc := redis.NewClient(&redis.Options{
		Addr:     c.RedisAddr,
		Password: "",
		DB:       0,
	})

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
	for {
		query := vc.FilterQuery{
			Addresses: []common.Address{
				l.contractAddr,
			},
			FromBlock: l.blockNumber,
		}

		logs, err := l.ec.FilterLogs(context.Background(), query)
		if err != nil {
			log.Fatal(err)
		}

		for _, vLog := range logs {
			if vLog.BlockNumber == l.blockNumber.Uint64() {
				time.Sleep(5 * time.Second)
				continue
			}

			l.blockNumber.SetUint64(vLog.BlockNumber)

			name := l.signatures[vLog.Topics[0].Hex()]
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
				l.logger.Infof("%s, %d, %s", "StreamRequested", v.StreamId.Uint64(), v.Client.Hex())

				pipelineId, userId, err := l.setCache(v.StreamId.Uint64())
				if err == nil {
					l.sendNotificationEvent(
						map[string]string{
							"name":        "StreamRequested",
							"pipeline_id": pipelineId,
							"user_id":     userId,
							"event":       fmt.Sprintf("stream id %s has been requested", v.StreamId.String()),
						},
					)
				} else {
					l.logger.Warn("StreamRequested set cache failed")
				}

			case *streamManager.ManagerStreamApproved:
				l.logger.Infof("%s, %d", "StreamApproved", v.StreamId.Uint64())

				pipelineId, userId, err := l.getCache(v.StreamId.Uint64())
				if err == nil {
					l.sendNotificationEvent(
						map[string]string{
							"name":        "StreamApproved",
							"pipeline_id": pipelineId,
							"user_id":     userId,
							"event":       fmt.Sprintf("stream id %s has been approved", v.StreamId.String()),
						},
					)
				} else {
					l.logger.Warn("StreamApproved get cache failed")
				}

			case *streamManager.ManagerStreamCreated:
				l.logger.Infof("%s, %d, %s", "StreamCreated", v.StreamId.Uint64(), v.StreamAddress)

				pipelineId, userId, err := l.getCache(v.StreamId.Uint64())
				if err == nil {
					l.sendNotificationEvent(
						map[string]string{
							"name":        "StreamCreated",
							"pipeline_id": pipelineId,
							"user_id":     userId,
							"event":       fmt.Sprintf("stream id %s has been created with stream address %s", v.StreamId.String(), v.StreamAddress.Hex()),
						},
					)
				} else {
					l.logger.Warn("StreamCreated get cache failed")
				}

			case *streamManager.ManagerStreamEnded:
				l.logger.Infof("%s, %d", "StreamEnded", v.StreamId.Uint64())

				pipelineId, userId, err := l.getCache(v.StreamId.Uint64())
				if err == nil {
					l.sendNotificationEvent(
						map[string]string{
							"name":        "StreamEnded",
							"pipeline_id": pipelineId,
							"user_id":     userId,
							"event":       fmt.Sprintf("stream id %s has been ended", v.StreamId.String()),
						},
					)
				} else {
					l.logger.Warn("StreamEnded get cache failed")
				}
			case *streamManager.ManagerInputChunkAdded:
				l.logger.Infof("%s, %d", "InputChunkAdded", v.StreamId.Uint64())

				pipelineId, userId, err := l.getCache(v.StreamId.Uint64())
				if err == nil {
					l.sendNotificationEvent(
						map[string]string{
							"name":        "InputChunkAdded",
							"pipeline_id": pipelineId,
							"user_id":     userId,
							"event":       fmt.Sprintf("input chunk has been added on stream id %s", v.StreamId.String()),
						},
					)
				} else {
					l.logger.Warn("InputChunkAdded get cache failed")
				}

			case *streamManager.ManagerRefundAllowed:
				l.logger.Infof("%s, %d", "RefundAllowed", v.StreamId.Uint64())

				pipelineId, userId, err := l.getCache(v.StreamId.Uint64())
				if err == nil {
					l.sendNotificationEvent(
						map[string]string{
							"name":        "RefundAllowed",
							"pipeline_id": pipelineId,
							"user_id":     userId,
							"event":       fmt.Sprintf("refund has been allowed on stream id %s", v.StreamId.String()),
						},
					)
				} else {
					l.logger.Warn("RefundAllowed get cache failed")
				}

			case *streamManager.ManagerRefundRevoked:
				l.logger.Infof("%s, %d", "RefundRevoked", v.StreamId.Uint64())

				pipelineId, userId, err := l.getCache(v.StreamId.Uint64())
				if err == nil {
					l.sendNotificationEvent(
						map[string]string{
							"name":        "RefundRevoked",
							"pipeline_id": pipelineId,
							"user_id":     userId,
							"event":       fmt.Sprintf("refund has been revoked on stream id %s", v.StreamId.String()),
						},
					)
				} else {
					l.logger.Warn("RefundRevoked get cache failed")
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
	p, err := l.ds.GetPipelineByStreamId(streamID)
	if err != nil {
		return "", "", err
	}

	err = l.rc.SetNX(fmt.Sprintf("u-%d", streamID), p.Id, 24*time.Hour).Err()
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
