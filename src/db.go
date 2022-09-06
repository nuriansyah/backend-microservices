package src

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func (p *Config) Close() {
	p.DB.Close()
}

func ConnectPostgres() (*Config, error) {
	connStr, err := loadPostgresConfig()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &Config{db}, nil
}
