package healthhndl

import (
	"net/http"

	"github.com/go-chi/render"
)

type Response struct {
	Status string `json:"status"`
}

func NewHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, Response{Status: "OK"})
	}
}
