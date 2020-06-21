package options

import (
	"hal9k/pkg/config"
	cliflag "k8s.io/component-base/cli/flag"
)

type Hal9000Options struct {
	QbotCfg    *config.QBotConfig
	ServiceCfg *config.ServiceConfig
}

func NewHal9000Options() *Hal9000Options {
	return &Hal9000Options{
		QbotCfg:    config.NewQBotConfig(),
		ServiceCfg: config.NewServiceConfig(),
	}
}

func (a *Hal9000Options) Flags() (fss cliflag.NamedFlagSets) {
	a.QbotCfg.AddFlags(fss.FlagSet("qbot"))
	a.ServiceCfg.AddFlags(fss.FlagSet("service"))
	return fss
}

func (a *Hal9000Options) Validate() []error {
	var errors []error
	errors = append(errors, a.QbotCfg.Validate()...)
	return errors
}
