package router

import (
	"linkshort/internal/transport/rest"

	"github.com/go-chi/chi/v5"
)

func NewRouter(handler *rest.Handlers) chi.Router {
	r := chi.NewRouter()

	r.Post("/", handler.CreateShortURL)
	r.Get("/{shortURL}", handler.GetFullUrl)

	return r

}
