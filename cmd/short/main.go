package main

import (
	"fmt"
	"linkshort/internal/consts"
	"linkshort/internal/transport/rest"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rest.CreateShortURL)
	mux.HandleFunc(fmt.Sprintf("/{%s}", consts.ShortURLParam), rest.GetFullUrl)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println(err)
	}
}
