package main

import "log"

func main() {
	configg := config{
		addr: ":8080",
	}
	app := &application{
		config: configg,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
