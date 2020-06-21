package config

import (
	"github.com/spf13/pflag"
	"hal9k/pkg/utils/reflectutils"
)

type ServiceConfig struct {
	Port     string `json:"port,omitempty" yaml:"port" description:"service port"`
	Loglevel string `json:"loglevel,omitempty" yaml:"loglevel" description:"service log level"`
}

func NewServiceConfig() *ServiceConfig {
	return &ServiceConfig{
		Port:     "8443",
		Loglevel: "info",
	}
}

func (q *ServiceConfig) Validate() []error {
	return nil
}

func (q *ServiceConfig) ApplyTo(cfg *ServiceConfig) {
	reflectutils.Override(cfg, q)
}

func (q *ServiceConfig) AddFlags(fs *pflag.FlagSet) {

	fs.StringVar(&q.Port, "port", q.Port, ""+
		"service listen addr port. default is 8443.")

	fs.StringVar(&q.Loglevel, "loglevel", q.Loglevel, ""+
		"Service log level. default is info.")
}
