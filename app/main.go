package main

import (
	"context"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/spf13/viper"

	"nossobr/article/delivery/http"
	"nossobr/article/delivery/http/middleware"
	repoArticle "nossobr/article/repository"
	articleUsecase "nossobr/article/usecase"
	repoAuthor "nossobr/author/repository"
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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := database.OpenConnections(); err != nil {
		log.Fatal("Unable to open connections to database: ", err)
	}
	defer database.CloseConnections()

	tx, err := database.NewTransaction(ctx, false)
	if err != nil {
		log.Fatal(err)
	}

	http.NewArticleHandler(e, articleUsecase.NewArticleUsecase(
		repoArticle.New(tx),
		repoAuthor.New(tx),
		time.Second*2,
	))

	log.Fatal(e.Start(viper.GetString("server.address"))) //nolint
}
