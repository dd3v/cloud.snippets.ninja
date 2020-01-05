package user

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type userHandler struct {
}

func NewHTTPHandler() *userHandler {
	return &userHandler{}
}

func (h *userHandler) Me(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(r.Header); err != nil {
		fmt.Println(err)
	}
}
