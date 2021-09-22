package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/mhdiiilham/segrato/config"
	"github.com/mhdiiilham/segrato/message"
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

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	api := app.Group("/api")
	v1 := api.Group("/v1")

	tokenService := token.TokenService{Config: &configuration}

	database := mongoDB.Database(configuration.Database)
	userCollection := database.Collection("user")
	userRepository := user.NewRepository(userCollection)
	userHandler := user.NewHandler(userRepository, tokenService)

	messageCollection := database.Collection("message")
	messageRepository := message.NewRepository(messageCollection)
	messageHandler := message.NewHandler(messageRepository)

	userRouter := v1.Group("/users")
	userRouter.Get(":userid", userHandler.GetUser)
	userRouter.Post("/", userHandler.RegisterUser)
	userRouter.Post("/login", userHandler.Login)
	userRouter.Get("/:id/messages", messageHandler.GetUserMessages)

	messageRouter := v1.Group("/messages")
	messageRouter.Post("/", messageHandler.PostMessage)
	messageRouter.Patch(":id", messageHandler.RepliedMessage)

	log.Fatal(app.Listen(":8081"))
}
