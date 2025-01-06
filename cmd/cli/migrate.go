package main

import (
	"log"
	"pujaprabha/internal/config"
	"pujaprabha/internal/domain/models"
)

func main() {
	db, err := config.ConfigDb()
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
	}
}
