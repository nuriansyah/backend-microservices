package main

import (
	"errors"
	"github.com/nuriansyah/log-mbkm-unpas/api"
	"github.com/nuriansyah/log-mbkm-unpas/repository"
	"github.com/nuriansyah/log-mbkm-unpas/src"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	//var dbPostgres *sql.DB
	//var err error

	dbPostgres, err := src.ConnectPostgres()
	if err != nil {
		panic(err)
	}
	if dbPostgres == nil {
		errors.New("Postgres not connection %v")
	}

	//SQLite
	//db, err := sql.Open("postgres", "file:basis-app.db")
	//if err != nil {
	//	panic(err)
	//}

	userRepo := repository.NewUserRepository(dbPostgres)
	//postsRepo := repository.NewPostRepository(dbPostgres)
	mainAPI := api.NewAPi(*userRepo)
	mainAPI.Start()
}
