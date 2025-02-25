package setting

import (
	"strings"
	"time"

	"github.com/spf13/viper"
)

var (
	Conf = &Configuration{}
)

type Configuration struct {
	RunMode               string        `env:"RUNMODE"`
	HTTPPort              string        `env:"HTTPPORT"`
	ReadTimeout           time.Duration `env:"READTIMEOUT"`
	WriteTimeout          time.Duration `env:"WRITETIMEOUT"`
	DefaultPageSize       int           `env:"DEFAULTPAGESIZE"`
	MaxPageSize           int           `env:"MAXPAGESIZE"`
	ServerShutdownTimeout time.Duration `env:"SERVERSHUTDOWNTIMEOUT"`
	LogSavePath           string        `env:"LOGSAVEPATH"`
	LogFileName           string        `env:"LOGFILENAME"`
	MaxSize               int           `env:"MAXSIZE"`
	MaxBackups            int           `env:"MAXBACKUPS"`
	Compress              bool          `env:"COMPRESS"`
	Level                 string        `env:"LEVEL"`
}

func Load(cfgFile string) error {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	err = viper.Unmarshal(Conf)
	if err != nil {
		return err
	}

	return nil
}
