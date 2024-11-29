package main

import (
	"fmt"
	"gracefullac/config"
	"gracefullac/pkg/database"
	"log"
)

func main() {
	fmt.Println("---- Implement Gracefull shutdown ----")

	var err error

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	_, err = database.NewDatabaseMysql(&cfg.Database)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("oke-----")

	// dbMysql.DB.Exec("")
}
