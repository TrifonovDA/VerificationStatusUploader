package database_tools

import (
	"context"
	"fmt"
	"github.com/deliryo-io/VerificationStatusUpdater/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

func NewConnection(ctx context.Context) *pgxpool.Pool {

	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v", config.BdCred.Username,
		config.BdCred.Password, config.BdCred.Host, config.BdCred.Port, config.BdCred.Database)

	dbpool, err := pgxpool.New(ctx, url)
	if err != nil {
		log.Printf("Unable to first connect to database_tools: %v\n", err)
	} else {
		fmt.Println("Database connection is good!")
	}

	return dbpool
}

func CloseConnectionPool(pool *pgxpool.Pool) {
	pool.Close()
}
