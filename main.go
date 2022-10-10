package main

import (
	"errors"
	"github.com/nuriansyah/log-mbkm-unpas/api"
	"github.com/nuriansyah/log-mbkm-unpas/repository"
	"github.com/nuriansyah/log-mbkm-unpas/src"
	"os"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASSWORD", "secret")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "1234")
	os.Setenv("DB_DATABASE", "log_km")
	dbPostgres, err := src.ConnectPostgres()
	if err != nil {
		panic(err)
	}
	if dbPostgres == nil {
		errors.New("Postgres not connection %v")
	}

	//SQLite
	//db, err := sql.Open("sqlite3", "file:basis-app.db")
	//if err != nil {
	//	panic(err)
	//}

	userRepo := repository.NewUserRepository(dbPostgres)
	postRepo := repository.NewPostRepository(dbPostgres)
	mainAPI := api.NewAPi(*userRepo, *postRepo)
	mainAPI.Start()
}
