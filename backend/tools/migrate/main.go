package main

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	// This is a dummy main function. The key is the imports above,
	// which register the postgres driver for the migrate CLI tool
	// when we build it.
	log.Println("Building custom migrate tool...")
}
