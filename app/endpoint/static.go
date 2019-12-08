package endpoint

import (
	"net/http"
)

func StaticEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("static endpoint"))
}
