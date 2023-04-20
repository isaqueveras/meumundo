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
	articleRepo "nossobr/article/repository/mongo"
	articleUsecase "nossobr/article/usecase"
	authorRepo "nossobr/author/repository/mongo"
	"nossobr/pkg/mongodb"
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
	e := echo.New()
	e.Use(middleware.InitMiddleware().CORS)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mongoConn, err := mongodb.NewMongoDBConn(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = mongoConn.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	authorRepo := authorRepo.NewMongoAuthorRepository(mongoConn)
	ar := articleRepo.NewMongoArticleRepository(mongoConn)
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	au := articleUsecase.NewArticleUsecase(ar, authorRepo, timeoutContext)
	http.NewArticleHandler(e, au)

	log.Fatal(e.Start(viper.GetString("server.address"))) //nolint
}
