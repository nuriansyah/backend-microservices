package main

import (
	"database/sql"
	"github.com/nuriansyah/log-mbkm-unpas/api"
	"github.com/nuriansyah/log-mbkm-unpas/repository"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "file:basis-app.db")
	if err != nil {
		panic(err)
	}

	userRepo := repository.NewUserRepository(db)
	postRepo := repository.NewPostRepository(db)

	mainAPI := api.NewAPi(*userRepo, *postRepo)
	mainAPI.Start()
}
