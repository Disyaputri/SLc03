package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"mini-soccer-cli/cli"
	"mini-soccer-cli/repository"
	"mini-soccer-cli/usecase"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := os.Getenv("MYSQL_DSN_CLI")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		log.Fatal("Gagal koneksi database: ", err)
	}

	fmt.Println("Database connected!")
	repo := repository.NewReportRepository(db)
	uc := usecase.NewReportUsecase(repo)
	cli.Run(context.Background(), uc)
}
