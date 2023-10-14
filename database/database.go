package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	"github.com/spf13/viper"
)

func OpenConnection() (*sql.DB, error) {
	var (
		dbHost = viper.GetString("database.host")
		dbPort = viper.GetString("database.port")
		dbUser = viper.GetString("database.user")
		dbPass = viper.GetString("database.pass")
		dbName = viper.GetString("database.name")

		dataSourceName = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, dbUser, dbName, dbPass)
		driverConfig   = stdlib.DriverConfig{
			ConnConfig: pgx.ConnConfig{
				RuntimeParams: map[string]string{
					"application_name": "meumundo",
					"DateStyle":        "ISO",
					"IntervalStyle":    "iso_8601",
					"search_path":      "public",
				},
			},
		}
	)

	stdlib.RegisterDriverConfig(&driverConfig)
	db, err := sql.Open("pgx", driverConfig.ConnectionString(dataSourceName))
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(120 * time.Second)
	db.SetMaxIdleConns(30)
	db.SetConnMaxIdleTime(20 * time.Second)
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
