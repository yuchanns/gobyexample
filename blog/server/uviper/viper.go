package uviper

import (
	"flag"
	"github.com/spf13/viper"
)

var path string

func init() {
	flag.StringVar(&path, "conf", "", "config path")
}

func Init() {
	viper.AddConfigPath(path)
}

func Get(fileName string) *Config {
	return &Config{fileName: fileName}
}

type Config struct {
	fileName string
}

func (c *Config) Unmarshal(data interface{}) error {
	viper.SetConfigName(c.fileName)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return viper.UnmarshalKey(c.fileName, data)
}
