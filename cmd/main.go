package main

import (
	"donezo/api/http"
	"donezo/app"
	"donezo/config"
	"flag"
	"log"
	"os"
)

var pathConfig string

func main() {
	cnf := loadConfig()
	app := app.NewMustApp(cnf)
	http.Run(cnf, &app)
}

func loadConfig() config.Config {
	const (
		defalutFlag = "$HOME/.donezo.yml"
		usage       = "pass config path"
	)

	flag.StringVar(&pathConfig, "c", defalutFlag, usage)
	flag.StringVar(&pathConfig, "config", defalutFlag, usage)
	flag.Parse()

	if envPath := os.Getenv("DONEZO_CONFIG_PATH"); len(envPath) > 0 {
		pathConfig = envPath
	}
	if len(pathConfig) == 0 {
		log.Fatal("configuration file not found")
	}

	cfg := config.MustReadConfig(pathConfig)
	return cfg
}
