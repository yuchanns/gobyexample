package uviper

import (
	"flag"
	"github.com/spf13/viper"
)

var path string

type Config struct {
	fileName string
}

func init() {
	flag.StringVar(&path, "conf", "", "config path")
}

func SetFlag(name string) {
	path = name
}

func Init() {
	if path == "" {
		panic("missing config path")
	}

	viper.AddConfigPath(path)
}

func Get(fileName string) *Config {
	return &Config{fileName: fileName}
}

func (c *Config) Unmarshal(data interface{}) (err error) {
	viper.SetConfigName(c.fileName)

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	err = viper.UnmarshalKey(c.fileName, data)

	return
}
