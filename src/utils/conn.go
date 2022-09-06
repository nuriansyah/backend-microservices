package utils

import (
	"database/sql"
	"github.com/nuriansyah/log-mbkm-unpas/src"
	"log"
)

type Conn struct {
	Postgres *sql.DB
}

func NewDBConn(cfg *src.Config) *Conn {
	return &Conn{Postgres: initPostgres(cfg)}
}

func initPostgres(cfg *src.Config) *sql.DB {
	db, err := sql.Open("postgres", cfg.DB.DSN)
	if err != nil {
		log.Fatal(err)
	}
	//db = db.Ping()
	//if err != nil {
	//	log.Fatal(err)
	//}
	return db
}
