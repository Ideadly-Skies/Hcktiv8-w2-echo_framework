// package config

// import (
// 	"context"
// 	"fmt"
// 	"log"

// 	_ "github.com/go-sql-driver/mysql"
// 	"github.com/jackc/pgx/v5/pgxpool"
// 	_ "github.com/jackc/pgx/v5/pgxpool"
// )

// var Pool *pgxpool.Pool

// func InitDB() {
// 	connStr := "user=postgres.goezqglikfdwpyvbafdy password=CristianoRonaldo13! host=aws-0-ap-southeast-1.pooler.supabase.com port=6543 dbname=postgres"

// 	config, err := pgxpool.Parseconfig(connStr)
// 	if err != nil {
// 		log.Fatalf("Failed to parsing config DB: %v", err)
// 	}

// 	Pool, err = pgxpool.NewWithConfig(context.Background(), config)
// 	if err != nil {
// 		log.Fatalf("failed to create db pooling: %v", err)
// 	}

// 	err = Pool.ping(context.Background(), config)
// 	if err != nil {
// 		log.Fatalf("db failed ping: %v", err)
// 	}

// 	fmt.Println("Database connected")
// }

// func CloseDB(){
// 	Pool.close()
// }

package config

import (
    "context"
    "fmt"
    "log"
    "time"

    "github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func InitDB() {
    connStr := "postgresql://postgres.goezqglikfdwpyvbafdy:CristianoRonaldo13!@aws-0-ap-southeast-1.pooler.supabase.com:6543/postgres"

    config, err := pgxpool.ParseConfig(connStr)
    if err != nil {
        log.Fatalf("Failed to parsing config DB: %v", err)
    }

    config.ConnConfig.ConnectTimeout = 5 * time.Second

    Pool, err = pgxpool.NewWithConfig(context.Background(), config)
    if err != nil {
        log.Fatalf("Failed to create db pooling: %v", err)
    }

    err = Pool.Ping(context.Background())
    if err != nil {
        log.Fatalf("DB failed Ping: %v", err)
    }

    fmt.Println("Database connected")
}

func CloseDB() {
    Pool.Close()
}