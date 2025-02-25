package cmd

import (
	"github.com/PainCodermax/to-do-list-api/pkg/logger"
	"github.com/PainCodermax/to-do-list-api/pkg/setting"
	"github.com/PainCodermax/to-do-list-api/pkg/signals"
	"github.com/PainCodermax/to-do-list-api/server"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	cfgFile string
)

func svrInit() error {
	if err := setting.Load(cfgFile); err != nil {
		return errors.Wrap(err, "loading config file failed")
	}
	gin.SetMode(setting.Conf.RunMode)
	logger.SetupLogger(setting.Conf)

	return nil
}

var rootCmd = &cobra.Command{
	Use:   "gin-clean-template",
	Short: "A clean architecture template for Golang Gin services",

	Run: func(_ *cobra.Command, _ []string) {
		if err := svrInit(); err != nil {
			log.Error("init server failed.", err)
			return
		}
		// start http server
		svr := server.NewServer(setting.Conf)
		if err := svr.Start(); err != nil {
			log.Error("init server failed.", err)
			return
		}

		// graceful shutdown
		stopCh := signals.SetupSignalHandler()
		sd, _ := signals.NewShutdown(setting.Conf.ServerShutdownTimeout)
		sd.Graceful(stopCh, svr)
	},
}

func Execute() error {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "./.env", "config file (default is ./.env)")
}
