package main

import (
	"log"
	"meumundo/database"
	"meumundo/delivery/http/estadual"
	"meumundo/delivery/http/federal"
	"meumundo/delivery/http/middleware"
	"meumundo/delivery/http/municipal"
	infraEstadual "meumundo/repository/estadual"
	infraFederal "meumundo/repository/federal"
	infraMunicipal "meumundo/repository/municipal"
	ucEstadual "meumundo/usecase/estadual"
	ucFederal "meumundo/usecase/federal"
	ucMunicipal "meumundo/usecase/municipal"
	"time"

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

	group := router.Group("v1")

	municipal.NewHandler(group, ucMunicipal.NewUsecase(infraEstadual.New(db), time.Second*2))
	estadual.NewHandler(group, ucEstadual.NewUsecase(infraMunicipal.New(db), time.Second*2))
	federal.NewHandler(group, ucFederal.NewUsecase(infraFederal.New(db), time.Second*2))

	log.Fatal(router.Start(viper.GetString("server.address")))
}
