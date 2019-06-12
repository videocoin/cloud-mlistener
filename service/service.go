package service

import (
	"github.com/VideoCoin/cloud-pkg/mqmux"
)

type Service struct {
	cfg      *Config
	eb       *EventBus
	listener *listener
}

func NewService(cfg *Config) (*Service, error) {
	mq, err := mqmux.NewWorkerMux(cfg.MQURI, cfg.Name)
	if err != nil {
		return nil, err
	}
	mq.Logger = cfg.Logger.WithField("system", "mq")

	eblogger := cfg.Logger.WithField("system", "eventbus")
	eb, err := NewEventBus(mq, eblogger)
	if err != nil {
		return nil, err
	}

	ds, err := NewDatastore(cfg.DBURI)
	if err != nil {
		return nil, err
	}

	elConfig := &ListenerConfig{
		NodeHTTPAddr: cfg.NodeHTTPAddr,
		ContractAddr: cfg.ContractAddr,
		RedisAddr:    cfg.RedisURI,
		EB:           eb,
		DS:           ds,
		Logger:       cfg.Logger,
	}

	listener, err := NewListener(elConfig)
	if err != nil {
		return nil, err
	}

	svc := &Service{
		cfg:      cfg,
		listener: listener,
		eb:       eb,
	}

	return svc, nil
}

func (s *Service) Start() error {
	go s.listener.Start()
	go s.eb.Start()
	return nil
}

func (s *Service) Stop() error {
	s.listener.Stop()
	s.eb.Stop()
	return nil
}
