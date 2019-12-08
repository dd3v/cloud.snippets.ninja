package app

import (
	"net/http"

	"github.com/dd3v/snippets.page.backend/app/endpoint"
	"github.com/dd3v/snippets.page.backend/app/storage"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type server struct {
	logger  *logrus.Logger
	router  *mux.Router
	storage storage.Storage
}

func newServer(storage storage.Storage) *server {
	server := &server{
		logger:  logrus.New(),
		router:  mux.NewRouter(),
		storage: storage,
	}
	server.initRouter()
	return server
}

func (server *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	server.router.ServeHTTP(w, r)
}

func (server *server) initRouter() {
	server.router.HandleFunc("/", endpoint.StaticEndpoint)
}
