package main

import (
	"donezo/config"
	"flag"
	"fmt"
	"log"
	"os"
)


var pathConfig string

func main() {
	cfg := loadConfig()
	fmt.Println(cfg)
}

func loadConfig() config.Config {
	const (
		defalutFlag = "$HOME/.donezo.yml"
		usage = "pass config path"
	)

	flag.StringVar(&pathConfig , "c" , defalutFlag , usage)
	flag.StringVar(&pathConfig , "config" , defalutFlag , usage)
	flag.Parse()

	if envPath := os.Getenv("DONEZO_CONFIG_PATH") ; len(envPath)>0 {
		pathConfig = envPath
	}
	if len(pathConfig) == 0 {
		log.Fatal("configuration file not found")
	}

	cfg := config.MustReadConfig(pathConfig)
	return cfg
}
