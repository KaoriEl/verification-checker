package command

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
)

var ctx = context.Background()

func DBConnection() *pgx.Conn {
	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return conn
}

func SelectUser() {

}
