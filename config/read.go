package config

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/viper"
)

func absPath(path string) (string, error) {
	if filepath.IsAbs(path) {
		return path, nil
	}
	return filepath.Abs(path)
}

func ReadConfig(path string) (Config, error) {
	var config Config
	abPath, err := absPath(path)
	if err != nil {
		return config, err
	}
	viper.SetConfigFile(abPath)
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return config, err
	}
	return config, viper.Unmarshal(&config)
}

func MustReadConfig(path string) Config {
	config, err := ReadConfig(path)
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	return config
}
