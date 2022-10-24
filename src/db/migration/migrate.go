package migration

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/nuriansyah/log-mbkm-unpas/src/db/seeder"
)

func Migrate(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY,
  	name varchar(255) NOT NULL,
 	email varchar(255) NOT NULL UNIQUE,
  	password varchar(255) NOT NULL,
  	role varchar(255) NOT NULL,
	avatar varchar(255) null
);
CREATE TABLE IF NOT EXISTS user_details (
    user_id INTEGER  NOT NULL,
  	nrp varchar(9) NOT NULL,
  	prodi varchar(255) NOT NULL,
  	program varchar(255) NOT NULL,
	company varchar(255) NULL,
	batch smallint NOT NULL,
	FOREIGN KEY (user_id) REFERENCES users(id)
);
CREATE TABLE IF NOT EXISTS posts (
    id INTEGER PRIMARY KEY,
  	author_id integer NOT NULL,
	title varchar(255) NOT NULL,
	description varchar(255) NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	FOREIGN KEY (author_id) REFERENCES users(id)
)
`)
	if err != nil {
		panic(err)
	}
	seeder.Seed(db)
}
