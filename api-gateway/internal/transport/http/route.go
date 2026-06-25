package httptr

import (
	"github.com/go-chi/chi/v5"

	healthhndl "github.com/uniqelus/opentrade/api-gateway/internal/transport/http/health"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", healthhndl.NewHandler())

	return r
}
