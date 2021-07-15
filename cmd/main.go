package main

import (
	"context"
	"fmt"
	"os"
	"post/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

var ctx = context.Background()
var db *pgxpool.Pool
var err error

func main() {
	p := os.Getenv("dbpass")

	db, err = pgxpool.Connect(ctx, "postgres://postgres:"+p+"@127.0.0.1/task")
	if err != nil {
		fmt.Println("connection to task failed")
	}

	a, err := storage.AllTasks(db)
	if err != nil {
		fmt.Println("Error")
	}

	fmt.Println(a)

	defer db.Close()
}
