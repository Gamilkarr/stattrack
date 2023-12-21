package main

import (
	"log"
	"net/http"

	config "github.com/Gamilkarr/stattrack/configs/server"
	"github.com/Gamilkarr/stattrack/internal/handlers"
	"github.com/Gamilkarr/stattrack/internal/repository"
)

func main() {
	repo, err := repository.NewRepo()
	if err != nil {
		log.Fatal(err)
	}
	e, err := handlers.NewHandler(repo)
	if err != nil {
		log.Fatal(err)
	}

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	r := e.NewRoute()

	log.Fatal(http.ListenAndServe(cfg.Address, r))
}
