package main

import (
	"errors"
	_ "github.com/lib/pq"
	"github.com/nuriansyah/log-mbkm-unpas/cmd/api"
	"github.com/nuriansyah/log-mbkm-unpas/src"
	repository2 "github.com/nuriansyah/log-mbkm-unpas/src/repository"
	"os"
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

	userRepo := repository2.NewUserRepository(dbPostgres)
	postRepo := repository2.NewPostRepository(dbPostgres)
	mainAPI := api.NewAPi(*userRepo, *postRepo)
	mainAPI.Start()
}
