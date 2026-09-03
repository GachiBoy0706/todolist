package main

import (
	"ToDoList/internal/core/tools"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dsn := os.Getenv("DATABASE_URL")
	fmt.Println("dsn: ", dsn)
	if dsn == "" {
		fmt.Println("no url for db")
		return
	}

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		fmt.Println("problems with pgxpool: ", err.Error())
		return
	}
	defer pool.Close()

	err = pool.Ping(ctx)
	if err != nil {
		fmt.Println("no ping: ", err.Error())
		return
	}
	fmt.Println("all good")

	router := mux.NewRouter()
	tools.Init(router, pool)

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
