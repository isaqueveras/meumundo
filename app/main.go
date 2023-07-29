package main

import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/spf13/viper"

	"nossobr/article/delivery/http"
	"nossobr/article/delivery/http/middleware"
	"nossobr/article/repository"
	"nossobr/article/usecase"
	"nossobr/database"
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

	e := echo.New()
	e.Use(middleware.InitMiddleware().CORS)

	db, err := database.OpenConnection()
	if err != nil {
		log.Fatal("Unable to open connections to database: ", err)
	}
	defer db.Close()

	http.NewArticleHandler(e, usecase.NewArticleUsecase(repository.NewRepo(db), time.Second*2))
	log.Fatal(e.Start(viper.GetString("server.address")))
}
