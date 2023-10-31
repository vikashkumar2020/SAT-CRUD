package main

import (
	"flag"
	"log"
	"os"
	config "sat-result/internal/config"
	db "sat-result/internal/infra/postgres"
	_ "sat-result/internal/infra/postgres/migrations"

	_ "github.com/jackc/pgx/v5"
	goose "github.com/pressly/goose/v3"
)

var (
	flags = flag.NewFlagSet("goose", flag.ExitOnError)
)

func main() {
	flags.Parse(os.Args[1:])
	args := flags.Args()

	if len(args) < 1 {
		flags.Usage()
		return
	}

	config := config.NewDBConfig()

	command := args[0]

	db, err := goose.OpenDBWithDriver("pgx", db.GetOrmConfig(config))
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v\n", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	arguments := []string{}
	if len(args) > 1 {
		arguments = append(arguments, args[1:]...)
	}

	if err := goose.Run(command, db, "./internal/infra/postgres/migrations", arguments...); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
}
