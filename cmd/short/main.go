package main

import (
	"fmt"
	"linkshort/internal/apllication"
	"linkshort/internal/repository/memory"
	"linkshort/internal/transport/rest"
	"linkshort/internal/transport/router"
	"net/http"
)

func main() {

	store := memory.NewStorage()
	app := apllication.NewApp(store)

	handler := rest.NewHandlers(app)

	rout := router.NewRouter(handler)

	if err := http.ListenAndServe(":8080", rout); err != nil {
		fmt.Println(err)
	}
}
