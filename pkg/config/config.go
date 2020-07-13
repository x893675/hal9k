package config

import (
	"fmt"
	"github.com/spf13/viper"
	"hal9k/pkg/logger"
)

const (
	DefaultConfigurationName = "config"
	DefaultConfigurationPath = "/etc/hal9000/"
)

var (
	// sharedConfig holds configuration
	sharedConfig *Config

	// shadowConfig contains options from commandline options
	shadowConfig = &Config{}
)

type Config struct {
	QBotCfg    *QBotConfig    `json:"qbot,omitempty" yaml:"qbot,omitempty" mapstructure:"qbot"`
	ServiceCfg *ServiceConfig `json:"service,omitempty" yaml:"service,omitempty" mapstructure:"service"`
}

func newConfig() *Config {
	return &Config{
		QBotCfg:    NewQBotConfig(),
		ServiceCfg: NewServiceConfig(),
	}
}

func Get() *Config {
	return sharedConfig
}

func (c *Config) Apply(conf *Config) {
	shadowConfig = conf

	if conf.QBotCfg != nil {
		conf.QBotCfg.ApplyTo(c.QBotCfg)
	}

	if conf.ServiceCfg != nil {
		conf.ServiceCfg.ApplyTo(c.ServiceCfg)
	}
}

func (c *Config) StripEmptyOptions() {
	if c.QBotCfg != nil && c.QBotCfg.QbotHttpEndpoint == "" {
		c.QBotCfg = nil
	}
}

// Load loads configuration after setup
func Load() error {
	sharedConfig = newConfig()

	viper.SetConfigName(DefaultConfigurationName)
	viper.AddConfigPath(DefaultConfigurationPath)
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logger.Warn(nil, "configuration file not found")
			return nil
		} else {
			panic(fmt.Errorf("error parsing configuration file %s", err))
		}
	}

	conf := newConfig()
	if err := viper.Unmarshal(conf); err != nil {
		logger.Error(nil, "error unmarshal configuration %v", err)
		return err
	} else {
		conf.Apply(shadowConfig)
		sharedConfig = conf
	}

	return nil
}
