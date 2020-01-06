package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"

	app "github.com/dd3v/snippets.page.backend/internal"
	"github.com/dd3v/snippets.page.backend/internal/config"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "../../config/app.toml", "path to config file")
}

func main() {
	config := config.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	if err := app.New(config); err != nil {
		log.Fatal(err)
	}

}
