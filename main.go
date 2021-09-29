package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/mhdiiilham/segrato/config"
	"github.com/mhdiiilham/segrato/message"
	"github.com/mhdiiilham/segrato/pkg/db"
	"github.com/mhdiiilham/segrato/pkg/token"
	service "github.com/mhdiiilham/segrato/services"
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

	app.Use(logger.New(logger.Config{
		Format:     "${pid} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "Asia/Jakarta",
	}))
	api := app.Group("/api")
	v1 := api.Group("/v1")

	tokenService := token.TokenService{Config: &configuration}

	database := mongoDB.Database(configuration.Database)
	userCollection := database.Collection("user")
	userRepository := user.NewRepository(userCollection)

	messageCollection := database.Collection("message")
	messageRepository := message.NewRepository(messageCollection)

	messageService := service.NewMessageService(messageRepository, userRepository)

	userHandler := user.NewHandler(userRepository, tokenService)
	messageHandler := message.NewHandler(messageService)

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
