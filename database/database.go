package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	"github.com/spf13/viper"
)

type database interface {
	// transaction uses a transaction from a connection already opened in the database
	transaction(ctx context.Context, readOnly bool) (interface{}, error)
	// Open open connection with database
	open() error
	// close connection with database
	close()
}

var connection database

// DBTransaction used to aggregate transactions
type DBTransaction struct {
	postgres *sql.Tx
	Builder  squirrel.StatementBuilderType
}

// OpenConnections open connections with database
func OpenConnections() (err error) {
	if connection != nil {
		return nil
	}

	connection = &postgres{}
	if err = connection.open(); err != nil {
		return err
	}
	return nil
}

// CloseConnections close all connections with database
func CloseConnections() {
	connection.close()
}

// NewTransaction uses a transaction from a connection already opened in the database
func NewTransaction(ctx context.Context, readOnly bool) (*DBTransaction, error) {
	tx := &DBTransaction{}

	pgsql, err := connection.transaction(ctx, readOnly)
	if err != nil {
		return nil, err
	}

	tx.postgres = pgsql.(*sql.Tx)
	tx.Builder = squirrel.StatementBuilder.
		PlaceholderFormat(squirrel.Dollar).
		RunWith(tx.postgres)

	return tx, nil
}

// Commit confirm pending transactions for all open databases
func (t *DBTransaction) Commit() (erro error) {
	return t.postgres.Commit()
}

// Rollback close all pending transaction for all open databases
func (t *DBTransaction) Rollback() {
	_ = t.postgres.Rollback()
}

// Query executes a query that returns rows, typically a SELECT.
func (t *DBTransaction) Query(query string, args ...any) (*sql.Rows, error) {
	return t.postgres.Query(query, args...)
}

// Execute executes a query that doesn't return rows, typically an INSERT/UPDATE/DELETE.
func (t *DBTransaction) Execute(query string, args ...any) (sql.Result, error) {
	return t.postgres.Exec(query, args...)
}

type postgres struct {
	db      *sql.DB
	timeout int
}

// open a transaction with the database
func (p *postgres) open() (err error) {
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
					"application_name": "nossobr",
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
		return err
	}

	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(120 * time.Second)
	db.SetMaxIdleConns(30)
	db.SetConnMaxIdleTime(20 * time.Second)
	if err = db.Ping(); err != nil {
		return err
	}

	p.db = db
	p.timeout = 2

	return
}

// close the connections with database
func (p *postgres) close() {
	if p.db != nil {
		_ = p.db.Close()
	}
}

// transaction opens a transaction on some already open connection
func (p *postgres) transaction(ctx context.Context, readOnly bool) (interface{}, error) {
	var (
		tx  *sql.Tx
		err error
	)

	ctx, cancel := context.WithCancel(ctx)
	go func() {
		<-time.After(time.Duration(p.timeout+1) * time.Second)
		if tx == nil {
			cancel()
		}
	}()

	if tx, err = p.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
		ReadOnly:  readOnly,
	}); err != nil {
		return nil, err
	}

	return tx, nil
}
