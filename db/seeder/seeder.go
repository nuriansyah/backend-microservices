package seeder

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"

	_ "github.com/mattn/go-sqlite3"
)

func Seed(db *sql.DB) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	rowUserMahasiswa, err := db.Exec("INSERT INTO users (name, email, password, role) VALUES ('Radit', 'resradit@gmail.com', ?, 'mahasiswa')", hashedPassword)
	if err != nil {
		panic(err)
	}

	userMahasiswaId, err := rowUserMahasiswa.LastInsertId()
	if err != nil {
		panic(err)
	}

	db.Exec("INSERT INTO user_details (user_id, program, company, batch) VALUES (?, 'MSIB', 'Binar Academy', 2019)", userMahasiswaId)

	// User Siswa
	rowUserDosen, err := db.Exec("INSERT INTO users (name, email, password, role) VALUES ('Dosen A', 'dosena@gmail.com', ?, 'dosen')", hashedPassword)
	if err != nil {
		panic(err)
	}

	userDosenId, err := rowUserDosen.LastInsertId()
	if err != nil {
		panic(err)
	}

	db.Exec("INSERT INTO user_details (user_id) VALUES (?)", userDosenId)

}
