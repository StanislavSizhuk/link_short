package main

import (
	"fmt"
	"linkshort/internal/apllication"
	"linkshort/internal/config"
	"linkshort/internal/repository/memory"
	"linkshort/internal/transport/rest"
	"linkshort/internal/transport/router"
	"net/http"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
		return
	}
	fmt.Printf("config: %s\n", cfg.Addr)
	store := memory.NewStorage()
	app := apllication.NewApp(store)

	handler := rest.NewHandlers(app)

	rout := router.NewRouter(handler)

	if err = http.ListenAndServe(cfg.Addr, rout); err != nil {
		fmt.Println(err)
	}
}
