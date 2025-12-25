package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chahar4/aura/adapter/handlers"
	"github.com/chahar4/aura/adapter/storages"
	"github.com/chahar4/aura/core/domains"
	"github.com/chahar4/aura/core/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load()
	dsn := "host=localhost user=root password=password dbname=auraDB port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	db.AutoMigrate(&domains.User{}, &domains.Role{}, &domains.Guild{}, &domains.GroupChannel{}, &domains.Channel{}, &domains.Message{})

	userRepository := storages.NewUserPostgresRepo(db)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	//auth
	r.Post("/register", userHandler.Register)
	r.Post("/login", userHandler.Login)
	r.Post("/forgotpassword", userHandler.ForgotPasswordSend)
	r.Post("/recovery", userHandler.ForgotPasswordRecovery)

	fmt.Print("server is up on port 3000")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err.Error())
	}
}
