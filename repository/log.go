package repository

import (
	"database/sql"
	"time"
)

type LogRepository struct {
	db *sql.DB
}

func NewLogRepository(db *sql.DB) *LogRepository {
	return &LogRepository{db: db}
}

type Log struct {
	ID        int       `db:"id"`
	Activity  string    `db:"activity"`
	CreatedAt time.Time `db:"created_at"`
}

func (l *LogRepository) InsertLog(mhs_id int, activity, created_at string) (int64, error) {
	sqlStmt := `INSERT INTO  log (mhs_id,log,created_at) VALUES (?,?,?);`

	tx, err := l.db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	result, err := tx.Exec(sqlStmt, mhs_id, activity, time.Now())
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	if err := tx.Commit(); err != nil {
		return 0, err
	}
	return id, err
}

func (l *LogRepository) FetchLogByID(logID, mhsID int) ([]Log, error) {
	sqlStmt := `SELECT * FROM log as l 
							WHERE EXISTS(
								SELECT m.id,m.name, FROM mahasiswa as m WHERE l.mhs_id = m.id)
				JOIN mahasiswa mhs ON l.mhs_id = mhs.id`
	tx, err := l.db.Begin()
	if err != nil {
		return nil, nil
	}
	defer tx.Rollback()

	rows, err := tx.Query(sqlStmt, logID, mhsID)

	defer rows.Close()
	var log []Log
	for rows.Next() {
		var logs Log
		err := rows.Scan(&logs.ID, &logs.Activity, &logs.CreatedAt)
		if err != nil {
			return nil, nil
		}
		log = append(log, logs)
	}
	if err := tx.Commit(); err != nil {
		return nil, nil
	}
	return log, nil
}
