package main

import (
	"database/sql"
	"github.com/nuriansyah/log-mbkm-unpas/api"
	"github.com/nuriansyah/log-mbkm-unpas/repository"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./basis-app.db")
	if err != nil {
		panic(err)
	}

	mhsRepo := repository.NewMahasiswaRepository(db)
	dosenRepo := repository.NewDosenRepository(db)
	mainAPI := api.NewAPi(*mhsRepo, *dosenRepo)
	mainAPI.Start()
}
