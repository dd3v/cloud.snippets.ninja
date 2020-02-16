package user

import (
	"context"
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

func (h handler) Me(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(r.Header); err != nil {
		fmt.Println(err)
	}
}

func (h handler) View(w http.ResponseWriter, r *http.Request) {

}

func (h handler) Create(w http.ResponseWriter, r *http.Request) {
	var request CreateRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {

	}
	err = request.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	q, _ := h.service.CreateUser(context.TODO(), request)

	fmt.Println(q)
	fmt.Printf("%v", request)

}
