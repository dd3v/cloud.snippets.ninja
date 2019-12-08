package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"

	"github.com/dd3v/snippets.page.backend/app"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "../config/app.toml", "path to config file")
}

func main() {
	config := app.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	app := app.NewApp(config)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
