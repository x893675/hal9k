package config

import (
	"github.com/spf13/pflag"
	"hal9k/pkg/utils/reflectutils"
)

type QBotConfig struct {
	WebhookPattern   string `json:"webhook_pattern,omitempty" yaml:"webhookPattern" description:"qbot webhook pattern"`
	Token            string `json:"token,omitempty" yaml:"token" description:"coolq http plugin token"`
	QbotHttpEndpoint string `json:"qbothttp_ep,omitempty" yaml:"qbothttpUrl" description:"coolq http plugin endpoint"`
	Secret           string `json:"secret,omitempty" yaml:"secret" description:"coolq http plugin secret"`
}

func NewQBotConfig() *QBotConfig {
	return &QBotConfig{
		WebhookPattern:   "/webhook_endpoint",
		Token:            "MyCoolqHttpToken",
		QbotHttpEndpoint: "ws://192.168.31.249:6700",
		Secret:           "CQHTTP_SECRET",
	}
}

func (q *QBotConfig) Validate() []error {
	return nil
}

func (q *QBotConfig) ApplyTo(cfg *QBotConfig) {
	reflectutils.Override(cfg, q)
}

func (q *QBotConfig) AddFlags(fs *pflag.FlagSet) {

	fs.StringVar(&q.WebhookPattern, "webhook-pattern", q.WebhookPattern, ""+
		"Database service host address. If left blank, the following related database options will be ignored.")

	fs.StringVar(&q.Token, "token", q.Token, ""+
		"Database service port number. If left blank, the following related database options will be ignored.")

	fs.StringVar(&q.QbotHttpEndpoint, "qbot-ep", q.QbotHttpEndpoint, ""+
		"Database database name. If left blank, the following related database options will be ignored.")

	fs.StringVar(&q.Secret, "secret", q.Secret, ""+
		"Username for access to database service.")
}
