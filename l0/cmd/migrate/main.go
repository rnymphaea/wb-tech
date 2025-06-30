package main

import (
	"flag"
	"log"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"l0/internal/config"
)

func main() {
	up := flag.Bool("up", false, "Run migrations up")
	down := flag.Bool("down", false, "Run migrations down")

	flag.Parse()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	databaseURL := cfg.GetDatabaseURL()
	m, err := migrate.New(
		"file://internal/database/migrations",
		databaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer m.Close()

	if *down {
		if err := m.Down(); err != nil {
			log.Fatal(err)
		}
		log.Println("migrations down completed")
	}
	if *up || (!*up && !*down) {
		if err := m.Up(); err != nil {
			log.Fatal(err)
		}
		log.Println("migrations up completed")
	}

}
