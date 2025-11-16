package main

import (
	"log"

	"github.com/chahar4/aura/db"
	"github.com/chahar4/aura/internal/repository"
)

func main() {
	db, err := db.NewDatabase()
	if err != nil {

		log.Fatalln(err.Error())
	}
	userRepo := repository.NewUserRepository(db.GetDB())
	userRepo.Init()
}
