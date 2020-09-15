package config

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

// NewConfig returns the Config structure initialized from the yaml file or environment
func NewConfig(configPath string) (config Config, err error) {
	v := viper.New()
	v.SetConfigType("yaml")
	v.AutomaticEnv()

	// default value
	v.SetDefault("database.host", "localhost")
	v.SetDefault("database.port", 3306)
	v.SetDefault("database.dbname", "article")
	v.SetDefault("database.user", "user")
	v.SetDefault("database.password", "password")
	v.SetDefault("server.port", 9090)
	v.SetDefault("context.timeout", "10s")

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// if config file not exist, use default value
	v.AddConfigPath(configPath)
	files, _ := ioutil.ReadDir(configPath)
	index := 0
	for _, file := range files {
		if filepath.Ext("./"+file.Name()) != ".yaml" && filepath.Ext("./"+file.Name()) != ".yml" {
			continue
		}

		v.SetConfigName(file.Name())
		var err error
		if index == 0 {
			err = v.ReadInConfig()
		} else {
			err = v.MergeInConfig()
		}
		if err == nil {
			index++
		}
	}

	if err = v.Unmarshal(&config); err != nil {
		return
	}

	return
}
