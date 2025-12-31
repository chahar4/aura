package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chahar4/aura/adapter/handlers"
	customMiddleware "github.com/chahar4/aura/adapter/middleware"
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
	db.AutoMigrate(&domains.User{}, &domains.Role{}, &domains.Guild{}, &domains.GuildMember{}, &domains.GroupChannel{}, &domains.Channel{}, &domains.Message{})

	//inject user
	userRepository := storages.NewUserPostgresRepo(db)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	//inject channel
	channelRepository := storages.NewChannelPostgresRepo(db)
	channelService := services.NewChannelService(channelRepository)
	channelHandler := handlers.NewChannelHandler(channelService)

	//inject guild member
	guildMemberRepository := storages.NewGuildMemberPostgresRepo(db)
	guildMemberService := services.NewGuildMemberService(guildMemberRepository)
	guildMemberHandler := handlers.NewGuildMemberHandler(guildMemberService)

	//inject guild
	guildRepository := storages.NewGuildPostgresRepo(db)
	guildService := services.NewGuildService(guildRepository, guildMemberRepository)
	guildHandler := handlers.NewGuildHandler(guildService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	authRoute := chi.NewRouter()
	authRoute.Use(customMiddleware.JwtMiddleware)

	authRoute.Post("/guilds", guildHandler.AddGuild)
	authRoute.Get("/users/me/guilds", guildMemberHandler.GetAllGuildsByUserID)

	authRoute.Post("/guilds/{id}/channels", channelHandler.AddChannel)
	authRoute.Get("/guilds/{id}/channels", channelHandler.GetAllChannelsByGroupChannelID)

	//auth
	r.Post("/register", userHandler.Register)
	r.Post("/login", userHandler.Login)
	r.Post("/forgotpassword", userHandler.ForgotPasswordSend)
	r.Post("/recovery", userHandler.ForgotPasswordRecovery)

	r.Mount("/", authRoute)

	fmt.Print("server is up on port 3000")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err.Error())
	}
}
