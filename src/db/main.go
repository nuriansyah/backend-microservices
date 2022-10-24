package main

import (
	"database/sql"
	"github.com/nuriansyah/log-mbkm-unpas/src/db/migration"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://localhost:5432/database?sslmode=disable")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	migration.Migrate(db)
}
