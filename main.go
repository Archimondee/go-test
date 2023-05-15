package main

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go-test/config"
	"go-test/controllers"
	"go-test/database"
	"go-test/database/migration"
	"go-test/routes"
	"go-test/services"
	"google.golang.org/api/option"
	"log"
	"os"
	"path/filepath"
)

var (
	ctx context.Context

	UserService         services.UserService
	UserController      controllers.UserController
	UserRouteController routes.UserRouteController

	AuthService         services.AuthService
	AuthController      controllers.AuthController
	AuthRouteController routes.AuthRouteController

	NotificationService         services.NotificationService
	NotificationController      controllers.NotificationController
	NotificationRouteController routes.NotificationRouteController

	TopicService         services.TopicService
	TopicController      controllers.TopicController
	TopicRouteController routes.TopicRouteController
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8000"
	} else {
		port = ":" + port
	}

	return port
}

func init() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.MySQLUser,
		config.MySQLPassword, config.MySQLHost, config.MySQLPort, config.MySQLDatabase)

	database.Database(url)
	migration.Migration()
}

func main() {
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	app.Use(compress.New())
	app.Use(cache.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, DELETE, PUT",
		AllowCredentials: true,
	}))

	ctx = context.TODO()
	Firebase, _, _ := SetupFirebase()

	UserService = services.NewUserServiceImpl(ctx)
	UserController = controllers.NewUserController(UserService)
	UserRouteController = routes.NewUserRouteController(UserController, UserService, NotificationController)

	AuthService = services.NewAuthService(ctx, UserService)
	AuthController = controllers.NewAuthController(AuthService, ctx)
	AuthRouteController = routes.NewAuthRouteController(AuthController)

	NotificationService = services.NewNotificationServiceImpl(ctx, UserService, TopicService)
	NotificationController = controllers.NewTokenController(UserService, NotificationService, Firebase)
	NotificationRouteController = routes.NewNotificationRouteontroller(NotificationController, UserService)

	TopicService = services.NewTopicServiceImpl(ctx)
	TopicController = controllers.NewTopicController(TopicService)
	TopicRouteController = routes.NewTopicRouteontroller(TopicController)

	AuthRouteController.AuthRoute(app)
	UserRouteController.UserRoute(app)
	NotificationRouteController.NotificationRoute(app)
	TopicRouteController.TopicRoute(app)

	//app.Listen(":8000")

	app.Listen(getPort())
}

func SetupFirebase() (*firebase.App, context.Context, *messaging.Client) {

	ctx := context.Background()

	serviceAccountKeyFilePath, err := filepath.Abs("./keys.json")
	if err != nil {
		panic("Unable to load serviceAccountKeys.json file")
	}

	opt := option.WithCredentialsFile(serviceAccountKeyFilePath)

	//Firebase admin SDK initialization
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic("Firebase load error")
	}

	//Messaging client
	client, _ := app.Messaging(ctx)

	return app, ctx, client
}
