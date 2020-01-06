package user

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type handler struct {
	service Service
}

func NewHTTPHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) Me(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(r.Header); err != nil {
		fmt.Println(err)
	}
}
