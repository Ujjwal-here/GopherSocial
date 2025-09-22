package main

import (
	"log"

	"github.com/Ujjwal-here/GopherSocial/internal/env"
	"github.com/Ujjwal-here/GopherSocial/internal/store"
)

func main() {
	configg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15min"),
		},
	}

	store := store.NewPostgresStorage(nil)

	app := &application{
		config: configg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
