package data

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

// DatabaseInterface describes database functions
type DatabaseInterface interface {
	SeattleOfficerMetadata() []map[string]string
	SeattleGetOfficerByBadge(badge string) (*SeattleOfficer, error)
	SeattleSearchOfficerByName(firstName, lastName string) ([]*SeattleOfficer, error)
	SeattleFuzzySearchByName(name string) ([]*SeattleOfficer, error)
	SeattleFuzzySearchByFirstName(firstName string) ([]*SeattleOfficer, error)
	SeattleFuzzySearchByLastName(lastName string) ([]*SeattleOfficer, error)

	TacomaOfficerMetadata() []map[string]string
	TacomaSearchOfficerByName(firstName, lastName string) ([]*TacomaOfficer, error)
	TacomaFuzzySearchByName(name string) ([]*TacomaOfficer, error)
	TacomaFuzzySearchByFirstName(firstName string) ([]*TacomaOfficer, error)
	TacomaFuzzySearchByLastName(lastName string) ([]*TacomaOfficer, error)
}

// Client is the client used to connect to the db
type Client struct {
	pool *pgxpool.Pool
}

// NewClient is the constructor for Client
func NewClient() *Client {
	var pool *pgxpool.Pool
	if os.Getenv("DATABASE_URL") != "" {
		pgxConfig, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
		if err != nil {
			log.Panic(fmt.Sprintf("Unable to create postgres connection config: %v", err))
		}
		pool, err = pgxpool.ConnectConfig(context.Background(), pgxConfig)
		if err != nil {
			log.Panic(fmt.Sprintf("Unable to create db connection: %v", err))
		}
	} else {
		// local set up
		username := os.Getenv("DB_USERNAME")
		password := os.Getenv("DB_PASSWORD")
		host := os.Getenv("DB_HOST")
		dbName := os.Getenv("DB_NAME")
		pgxConfig, err := pgxpool.ParseConfig(fmt.Sprintf("postgresql://%s", host))
		if err != nil {
			log.Panic(fmt.Sprintf("Unable to create postgres connection config: %v", err))
		}

		pgxConfig.ConnConfig.Port = 5432
		pgxConfig.ConnConfig.Database = dbName
		pgxConfig.ConnConfig.User = username
		pgxConfig.ConnConfig.Password = password

		pool, err = pgxpool.ConnectConfig(context.Background(), pgxConfig)
		if err != nil {
			log.Panic(fmt.Sprintf("Unable to create db connection: %v", err))
		}
	}

	return &Client{pool: pool}
}
