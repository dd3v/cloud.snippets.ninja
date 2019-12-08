package app

import (
	"net/http"

	"github.com/dd3v/snippets.page.backend/app/endpoint"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

//App default app struct
type App struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

//NewApp create new app instance
func NewApp(config *Config) *App {
	return &App{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

//Run - application configuration and launch
func (app *App) Run() error {
	if err := app.setLogLevel(); err != nil {
		return err
	}
	app.setRouter()
	app.logger.Info("server is started...")

	if err := http.ListenAndServe(app.config.BindAddr, app.router); err != nil {
		app.logger.Error(err)
	}

	return nil
}

func (app *App) setRouter() {
	app.router.HandleFunc("/", endpoint.StaticEndpoint)
}

func (app *App) setLogLevel() error {
	level, err := logrus.ParseLevel(app.config.LogLevel)
	if err != nil {
		return err
	}
	app.logger.SetLevel(level)

	return nil
}
