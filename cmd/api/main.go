package main

import (
	"log"

	"github.com/Ujjwal-here/GopherSocial/internal/env"
)

func main() {
	configg := config{
		addr: env.GetString("ADDR", ":8080"),
	}
	app := &application{
		config: configg,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
