package main

import (
	"log"
	"time"

	"nossobr/database"
	http "nossobr/delivery/http/article"
	"nossobr/delivery/http/middleware"
	"nossobr/repository/article"
	"nossobr/usecase"

	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if viper.GetBool("debug") {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	router := echo.New()
	router.Use(middleware.InitMiddleware().CORS)

	db, err := database.OpenConnection()
	if err != nil {
		log.Fatal("Unable to open connections to database: ", err)
	}
	defer db.Close()

	articleRepo := article.New(db)
	uc := usecase.NewArticleUsecase(articleRepo, time.Second)

	articleGroup := router.Group("article")
	http.NewArticleHandler(articleGroup, uc)

	log.Fatal(router.Start(viper.GetString("server.address")))
}
