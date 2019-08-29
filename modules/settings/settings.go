package settings

import (
	"os"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	Debug bool
	Port  string
)

const (
	configName   = "config"
	configPath   = "."
	configFormat = "yaml"
)

func Init() {
	logrus.SetOutput(os.Stdout)

	v := get()
	v.SetConfigType(configFormat)
	v.AddConfigPath(configPath)
	v.SetConfigName(configName)

	if err := v.ReadInConfig(); err != nil {
		logrus.Panicln("load config file failed", err)
	}

	Debug = v.GetBool("app.debug")
	if Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	Port = v.GetString("app.port")
}

var (
	once          sync.Once
	viperInstance *viper.Viper
)

func get() *viper.Viper {
	once.Do(func() {
		viperInstance = viper.New()
	})
	return viperInstance
}
