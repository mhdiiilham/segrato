package main

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mhdiiilham/segrato/config"
	"github.com/mhdiiilham/segrato/pkg/db"
	"github.com/mhdiiilham/segrato/pkg/token"
	"github.com/mhdiiilham/segrato/user"
)

func main() {
	fmt.Println("Starting SEGRATO API")

	configuration, configErr := config.ReadConfig()
	if configErr != nil {
		panic(configErr)
	}

	fmt.Println("Production Mode Enabled:", strconv.FormatBool(configuration.Production))

	mongoDB, err := db.NewMongoDBConnection(configuration.MongoDBURI)
	if err != nil {
		panic(err)
	}

	app := fiber.New()
	api := app.Group("/api")
	v1 := api.Group("/v1")

	tokenService := token.TokenService{Config: &configuration}

	database := mongoDB.Database(configuration.Database)
	userCollection := database.Collection("user")
	userRepository := user.NewRepository(userCollection)
	userHandler := user.NewHandler(userRepository, tokenService)

	userRouter := v1.Group("/users")
	userRouter.Post("/", userHandler.RegisterUser)
	userRouter.Post("/login", userHandler.Login)

	app.Listen(":8080")

}
