package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/spf13/viper"

	"nossobr/article/delivery/http"
	"nossobr/article/delivery/http/middleware"
	articleRepo "nossobr/article/repository/mysql"
	articleUsecase "nossobr/article/usecase"
	authorRepo "nossobr/author/repository/mysql"
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
	dbHost := viper.GetString("database.host")
	dbPort := viper.GetString("database.port")
	dbUser := viper.GetString("database.user")
	dbPass := viper.GetString("database.pass")
	dbName := viper.GetString("database.name")
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "America/Fortaleza")

	dbConn, err := sql.Open("mysql", fmt.Sprintf("%s?%s", connection, val.Encode()))
	if err != nil {
		log.Fatal(err)
	}

	if err = dbConn.Ping(); err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = dbConn.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	e := echo.New()
	e.Use(middleware.InitMiddleware().CORS)

	authorRepo := authorRepo.NewMysqlAuthorRepository(dbConn)
	ar := articleRepo.NewMysqlArticleRepository(dbConn)
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	au := articleUsecase.NewArticleUsecase(ar, authorRepo, timeoutContext)
	http.NewArticleHandler(e, au)

	log.Fatal(e.Start(viper.GetString("server.address"))) //nolint
}
