package app

import (
	"github.com/spf13/cobra"
	"hal9k/cmd/app/options"
	"hal9k/internal"
	"hal9k/pkg/client/qbot"
	serverconfig "hal9k/pkg/config"
	"hal9k/pkg/logger"
	"hal9k/pkg/utils/signals"
	"hal9k/pkg/version"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
)

func NewHal9000Command() *cobra.Command {
	s := options.NewHal9000Options()

	cmd := &cobra.Command{
		Use:  "account-service",
		Long: `account service`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := serverconfig.Load()
			if err != nil {
				return err
			}

			err = Complete(s)
			if err != nil {
				return err
			}

			if errs := s.Validate(); len(errs) != 0 {
				return utilerrors.NewAggregate(errs)
			}

			return Run(s, signals.SetupSignalHandler())
		},
	}

	fs := cmd.Flags()
	namedFlagSets := s.Flags()

	for _, f := range namedFlagSets.FlagSets {
		fs.AddFlagSet(f)
	}

	return cmd
}

// apply server run options to configuration
func Complete(s *options.Hal9000Options) error {

	// loading configuration file
	conf := serverconfig.Get()

	conf.Apply(&serverconfig.Config{
		ServiceCfg: s.ServiceCfg,
		QBotCfg:    s.QbotCfg,
	})

	*s = options.Hal9000Options{
		ServiceCfg: conf.ServiceCfg,
		QbotCfg:    conf.QBotCfg,
	}

	return nil
}

func Run(s *options.Hal9000Options, stopCh <-chan struct{}) error {
	logger.SetLevelByString(s.ServiceCfg.Loglevel)

	err := qbot.NewQbot(s.QbotCfg, stopCh)
	if err != nil {
		return err
	}
	logger.Info(nil, "service verson is %s", version.Version)
	return internal.NewHttpServer(s.ServiceCfg)
}
