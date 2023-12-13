package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	// These module inits register the postgres and file drivers
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("%s <path> <up|down|force>\n", os.Args[0])
	}
	dbHost := os.Getenv("DBHOST")
	dbPassword := os.Getenv("DBPASSWORD")
	dbUser := os.Getenv("DBUSER")
	dbName := os.Getenv("DATABASE")
	dbPort := os.Getenv("DBPORT")
	dbSsl := os.Getenv("DBSSL")

	if len(dbPort) == 0 {
		dbPort = "5432"
	}

	if len(dbSsl) == 0 {
		dbSsl = "require"
	}

	path := os.Args[1]
	cmd := os.Args[2]

	sourcePath := fmt.Sprintf("file://%s", path)
	databaseUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		url.QueryEscape(dbUser),
		url.QueryEscape(dbPassword),
		dbHost,
		dbPort,
		dbName,
		dbSsl)

	m, err := migrate.New(
		sourcePath,
		databaseUrl)
	if err != nil {
		log.Fatalln(err)
	}
	switch cmd {
	case "up":
		if !handleMigrateError(m.Up()) {
			return
		}
	case "down":
		if !handleMigrateError(m.Down()) {
			return
		}
	case "step":
		if len(os.Args) < 4 {
			log.Fatalf("Need a step number (positive or negative)")
		}
		if !handleMigrateError(m.Steps(mustReadInt(os.Args[3]))) {
			return
		}
	case "force":
		if len(os.Args) < 4 {
			log.Fatalf("Need a force version")
		}
		if !handleMigrateError(m.Force(mustReadInt(os.Args[3]))) {
			return
		}
	default:
		log.Fatalln("command does not exist")
	}
	log.Printf("Migration of DB [%s] for host [%s:%s] successful\n", dbName, dbHost, dbPort)
}
func handleMigrateError(err error) bool {
	if err != nil && err != migrate.ErrNoChange {
		log.Fatalln(err)
		return false
	} else if err == migrate.ErrNoChange {
		log.Println(err)
		return true
	}
	return true
}

func mustReadInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalln(err)
	}
	return i
}
