package service

import (
	notificationv1 "github.com/VideoCoin/cloud-api/notifications/v1"
	"github.com/VideoCoin/cloud-pkg/mqmux"
	"github.com/sirupsen/logrus"
)

type EventBus struct {
	logger *logrus.Entry
	mq     *mqmux.WorkerMux
}

func NewEventBus(mq *mqmux.WorkerMux, logger *logrus.Entry) (*EventBus, error) {
	return &EventBus{
		logger: logger,
		mq:     mq,
	}, nil
}

func (e *EventBus) Start() error {
	if err := e.registerPublishers(); err != nil {
		return err
	}

	if err := e.registerConsumers(); err != nil {
		return err
	}

	return e.mq.Run()
}

func (e *EventBus) Stop() error {
	return e.mq.Close()
}

func (e *EventBus) registerPublishers() error {
	if err := e.mq.Publisher("listen/stream"); err != nil {
		return err
	}

	if err := e.mq.Publisher("notifications/send"); err != nil {
		return err
	}

	return nil
}

func (e *EventBus) registerConsumers() error {
	return nil
}

func (e *EventBus) SendNotification(req *notificationv1.Notification) error {
	e.logger.Infof("sending notification %v", req)
	return e.mq.Publish("notifications/send", req)
}
