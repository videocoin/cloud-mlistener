package service

import (
	eventsv1 "github.com/VideoCoin/cloud-api/events/v1"
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
	err := e.registerPublishers()
	if err != nil {
		return err
	}

	err = e.registerConsumers()
	if err != nil {
		return err
	}

	return e.mq.Run()
}

func (e *EventBus) Stop() error {
	return e.mq.Close()
}

func (e *EventBus) registerPublishers() error {
	err := e.mq.Publisher("events/new")
	if err != nil {
		return err
	}

	err = e.mq.Publisher("listen/stream")
	if err != nil {
		return err
	}

	return nil
}

func (e *EventBus) registerConsumers() error {
	return nil
}

func (e *EventBus) CreateEvent(req *eventsv1.Event) error {
	return e.mq.Publish("events/new", req)
}

func (e *EventBus) CreateStreamListener(req *eventsv1.Event) error {
	return e.mq.Publish("listen/stream", req)
}

func (e *EventBus) SendNotification(req *notificationv1.Notification) error {
	return e.mq.Publish("notifications/send", req)
}
