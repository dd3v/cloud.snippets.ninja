package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/BurntSushi/toml"

	app "github.com/dd3v/snippets.page.backend/internal"
	"github.com/dd3v/snippets.page.backend/internal/config"
	"github.com/dd3v/snippets.page.backend/internal/user"
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

	address := fmt.Sprintf(":%v", config.BindAddr)
	httpServer := &http.Server{
		Addr:    address,
		Handler: buildRouter(),
	}

	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Println(err)
		os.Exit(-1)
	}

}

func buildRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	v1 := router.PathPrefix("/v1").Subrouter()

	userHandler := user.NewHTTPHandler()

	v1.HandleFunc("/me", userHandler.Me).Methods("GET")

	return router
}
