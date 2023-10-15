package main

import (
	"log"
	"time"

	"meumundo/database"
	"meumundo/delivery/http"
	"meumundo/delivery/http/middleware"

	infraEstadual "meumundo/repository/estadual"
	infraFederal "meumundo/repository/federal"
	infraMunicipal "meumundo/repository/municipal"

	ucEstadual "meumundo/usecase/estadual"
	ucFederal "meumundo/usecase/federal"
	ucMunicipal "meumundo/usecase/municipal"

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

	repoMunicipal := infraMunicipal.New(db)
	repoEstadual := infraEstadual.New(db)
	repoFederal := infraFederal.New(db)

	usecaseMunicipal := ucEstadual.NewUsecase(repoMunicipal, time.Second)
	usecaseEstadual := ucMunicipal.NewUsecase(repoEstadual, time.Second)
	usecaseFederal := ucFederal.NewUsecase(repoFederal, time.Second)

	group := router.Group("v1")
	http.NewHandler(group, uc)

	log.Fatal(router.Start(viper.GetString("server.address")))
}
