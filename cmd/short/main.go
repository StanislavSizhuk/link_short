package main

import (
	"fmt"
	"linkshort/internal/apllication"
	"linkshort/internal/consts"
	"linkshort/internal/repository/memory"
	"linkshort/internal/transport/rest"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	store := memory.NewStorage()
	app := apllication.NewApp(store)

	handler := rest.NewHandlers(app)

	mux.HandleFunc("/", handler.CreateShortURL)

	mux.HandleFunc(fmt.Sprintf("/{%s}", consts.ShortURLParam), handler.GetFullUrl)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println(err)
	}
}
