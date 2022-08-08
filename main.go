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

	mhsRepo := repository.NewMahasiswaRepository(db)
	dosenRepo := repository.NewDosenRepository(db)
	logRepo := repository.NewLogRepository(db)
	mainAPI := api.NewAPi(*mhsRepo, *dosenRepo, *logRepo)
	mainAPI.Start()
}
