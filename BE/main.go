package main

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/Hemtrakan/todolist_connection_mongodb/docs"
	"github.com/Hemtrakan/todolist_connection_mongodb/pkg/api"
	"github.com/Hemtrakan/todolist_connection_mongodb/pkg/api/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @securityDefinitions.apikey Bearer
// @in 		header
// @name 	Authorization
func main() {
	fmt.Println("start api")
	initConfig()
	initTimeZone()
	db, client, cancel, ctx := initDatabase()
	defer cancel()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	router := gin.New()
	if viper.GetString("app.environment") == "local" {
		router = gin.Default()
	}

	gin.SetMode(gin.ReleaseMode)
	router.Use(gin.Recovery())
	router.Use(middlewares.TrackingId())
	router.Use(middlewares.CORSMiddleware())
	router.Use(middlewares.TimeoutMiddleware(60 * time.Second))

	api.New(router, db, client)
	if viper.GetBool("app.isSwagger") {
		initSwagger()
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", viper.GetInt("app.port")),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()
	fmt.Println("start api success")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

}

func initConfig() {
	fmt.Println("initConfig")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	// viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initTimeZone() {
	fmt.Println("initTimeZone")
	timezone, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}
	time.Local = timezone
}

func initSwagger() {
	fmt.Println("initSwagger")
	docs.SwaggerInfo.Title = viper.GetString("app.name")
	docs.SwaggerInfo.Description = "Environment : " + viper.GetString("app.environment")
	docs.SwaggerInfo.Version = "version : " + viper.GetString("app.version")
	docs.SwaggerInfo.Host = viper.GetString("app.host") + "/"
	docs.SwaggerInfo.BasePath = viper.GetString("app.path")
}

func initDatabase() (*mongo.Database, *mongo.Client, context.CancelFunc, context.Context) {
	fmt.Println("initDatabase")
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	connectionString := viper.GetString("mongodb.connection")

	mongoOptions := options.Client().ApplyURI(connectionString)
	mongoOptions.SetMaxPoolSize(50)

	client, err := mongo.Connect(ctx, mongoOptions)
	if err != nil {
		log.Fatalf("database connect fail %v!!", err)
		cancel()
	}
	db := client.Database(viper.GetString("mongodb.database"))

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("database ping fail %v!!", err)
		cancel()
	}

	fmt.Println("Connected to MongoDB!")
	return db, client, cancel, ctx
}
