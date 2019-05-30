package service

import (
	"github.com/sirupsen/logrus"
)

type Config struct {
	Name    string `envconfig:"-"`
	Version string `envconfig:"-"`

	NodeHTTPAddr string `default:"" envconfig:"NODEHTTPADDR"`
	ContractAddr string `default:"0x8bFEF9CaB41460dB146b1E701B782edBa7883bC9" envconfig:"CONTRACTADDR"`

	DBURI    string `default:"root:@tcp(127.0.0.1:3306)/videocoin?charset=utf8&parseTime=True&loc=Local" envconfig:"DBURI"`
	RedisURI string `default:"" envconfig:"REDISURI"`
	MQURI    string `default:"amqp://guest:guest@127.0.0.1:5672" envconfig:"MQURI"`

	Logger *logrus.Entry `envconfig:"-"`
}
