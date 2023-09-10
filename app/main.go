package main

import (
	"log"
	"time"

	"nossobr/database"
	"nossobr/delivery/http"
	"nossobr/delivery/http/middleware"
	"nossobr/repository"
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

	repo := repository.New(db)
	uc := usecase.NewUsecase(repo, time.Second)

	group := router.Group("v1")
	http.NewHandler(group, uc)

	log.Fatal(router.Start(viper.GetString("server.address")))
}
