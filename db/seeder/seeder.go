package seeder

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
)

func Seed(db *sql.DB) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	_, err := db.Exec("INSERT INTO mahasiswa (id,name, email, password,dosen_id) VALUES (1,'Radit', 'resradit@gmail.com', ?, 'mahasiswa',$1)", hashedPassword)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("INSERT INTO dosen (id,name, email, password) VALUES (1,'Dadang', 'dadang@gmail.com', ?, 'mahasiswa')", hashedPassword)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("INSERT INTO log (mhs_,id,activity,created_at) VALUES ($2,`activity day1`,datetime(`now`)) ")
	if err != nil {
		panic(err)
	}

}
