package rest

import (
	"fmt"
	"io"
	"linkshort/internal/apllication"
	"linkshort/internal/consts"
	"linkshort/internal/domain/url"

	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handlers struct {
	app *apllication.App
}

func NewHandlers(app *apllication.App) *Handlers {
	return &Handlers{app: app}
}

func (h *Handlers) CreateShortURL(w http.ResponseWriter, r *http.Request) {

	var (
		fullUrlbytes []byte
		fullUrl      string
		modelURL     *url.URL
		err          error
	)
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	fullUrlbytes, err = io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fullUrl = string(fullUrlbytes)
	modelURL, err = h.app.CreateShortURL(r.Context(), fullUrl)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = fmt.Fprintf(w, "http://localhost:8080/%s", modelURL.Short())

}

func (h *Handlers) GetFullUrl(w http.ResponseWriter, r *http.Request) {
	var (
		shortURL string
		err      error
		modelURL *url.URL
	)

	shortURL = chi.URLParam(r, consts.ShortURLParam)

	modelURL, err = h.app.GetFullURL(r.Context(), shortURL)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	w.Header().Set("Location", modelURL.Full())
	w.WriteHeader(http.StatusTemporaryRedirect)

}
