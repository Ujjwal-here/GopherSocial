package main

import (
	"log"

	"github.com/Ujjwal-here/GopherSocial/internal/db"
	"github.com/Ujjwal-here/GopherSocial/internal/env"
	"github.com/Ujjwal-here/GopherSocial/internal/store"
)

func main() {
	configg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	db, err := db.New(configg.db.addr, configg.db.maxOpenConns, configg.db.maxIdleConns, configg.db.maxIdleTime)

	if err != nil {
		log.Panic(err)
	}

	store := store.NewPostgresStorage(db)

	app := &application{
		config: configg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
