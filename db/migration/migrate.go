package migration

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nuriansyah/log-mbkm-unpas/db/seeder"
)

func Migrate(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS dosen (
    id integer not null primary key AUTOINCREMENT,
    name varchar(255) not null,
    email varchar(255) not null UNIQUE,
    password varchar(255) not null,
	avatar varchar(255) null
);
CREATE TABLE IF NOT EXISTS mahasiswa (
    id integer not null primary key AUTOINCREMENT,
    name varchar(255) not null,
    email varchar(255) not null UNIQUE,
    password varchar(255) not null,
	avatar varchar(255) null
);
CREATE TABLE IF NOT EXISTS log (
    id integer not null primary key AUTOINCREMENT,
	log varchar(255) not null,
    sqltime TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
)
`)
	if err != nil {
		panic(err)
	}
	seeder.Seed(db)
}
