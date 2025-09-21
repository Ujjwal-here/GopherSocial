package main

import (
	"log"

	"github.com/Ujjwal-here/GopherSocial/internal/env"
	"github.com/Ujjwal-here/GopherSocial/internal/store"
)

func main() {
	configg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	store := store.NewPostgresStorage(nil)

	app := &application{
		config: configg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
