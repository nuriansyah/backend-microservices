package migration

import (
	"database/sql"
	"github.com/nuriansyah/log-mbkm-unpas/db/seeder"

	_ "github.com/mattn/go-sqlite3"
)

func Migrate(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
    id integer not null primary key AUTOINCREMENT,
    name varchar(255) not null,
    email varchar(255) not null UNIQUE,
    password varchar(255) not null,
	role varchar(255) not null,
	avatar varchar(255) null
);
CREATE TABLE IF NOT EXISTS user_details (
    user_id integer NOT NULL,
	program varchar(255) NOT NULL,
	company varchar(255) NULL,
	batch smallint UNSIGNED NULL,
	FOREIGN KEY (user_id) REFERENCES users(id)
);
CREATE TABLE IF NOT EXISTS posts(
	id integer NOT NULL PRIMARY KEY AUTOINCREMENT,
	author_id integer NOT NULL,
	title varchar(255) NOT NULL,
	desc text NOT NULL,
	created_at datetime NOT NULL,
	FOREIGN KEY (author_id) REFERENCES users(id)
);
`)
	if err != nil {
		panic(err)
	}

	seeder.Seed(db)
}
