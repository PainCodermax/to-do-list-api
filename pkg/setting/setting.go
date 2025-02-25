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
	RunMode               string        `env:"RunMode"`
	HTTPPort              string        `env:"HTTPPort"`
	ReadTimeout           time.Duration `env:"ReadTimeout"`
	WriteTimeout          time.Duration `env:"WriteTimeout"`
	DefaultPageSize       int           `env:"DefaultPageSize"`
	MaxPageSize           int           `env:"MaxPageSize"`
	ServerShutdownTimeout time.Duration `env:"ServerShutdownTimeout"`
	LogSavePath           string        `env:"LogSavePath"`
	LogFileName           string        `env:"LogFileName"`
	MaxSize               int           `env:"MaxSize"`
	MaxBackups            int           `env:"MaxBackups"`
	Compress              bool          `env:"Compress"`
	Level                 string        `env:"Level"`
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
