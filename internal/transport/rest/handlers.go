package rest

import (
	"fmt"
	"io"
	"linkshort/internal/apllication"
	"linkshort/internal/consts"
	"linkshort/internal/domain/url"

	"net/http"
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
	)

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	shortURL = r.PathValue(consts.ShortURLParam)
	fmt.Println("shortURL:", shortURL)
	w.WriteHeader(http.StatusNotFound)
}
