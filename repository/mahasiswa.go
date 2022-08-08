package repository

import (
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type MahasiswaRepository struct {
	db *sql.DB
}

func NewMahasiswaRepository(db *sql.DB) *MahasiswaRepository {
	return &MahasiswaRepository{db: db}
}

func (u *MahasiswaRepository) Login(email string, password string) (*int, error) {
	statement := "SELECT id,email,password,is_logged_in  FROM users WHERE email = ?"
	res := u.db.QueryRow(statement, email, password)
	var hashedPassword string
	var id int
	res.Scan(&id, &hashedPassword)
	if bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) != nil {
		return nil, errors.New("Login Failed")
	}
	statement = "UPDATE user SET is_logged_in = 1 WHERE id = ?"

	_, err := u.db.Exec(statement, id)
	if err != nil {
		return nil, err
	}
	return &id, nil
}
func (u *MahasiswaRepository) GetUserData(id int) (*Mahasiswa, error) {
	statement := "SELECT id,name,email,avatar FROM mahasiswa "
	var mhs Mahasiswa
	res := u.db.QueryRow(statement, id)
	err := res.Scan(&mhs.Id, &mhs.Name, &mhs.Email, &mhs.Avatar)
	return &mhs, err
}
func (u *MahasiswaRepository) CheckEmail(email string) (bool, error) {
	statement := "SELECT count(*) FROM mahasiswa WHERE email = ?"
	res := u.db.QueryRow(statement, email)

	var count int
	err := res.Scan(&count)
	if count > 0 {
		return false, err
	}
	return true, err
}
func (u *MahasiswaRepository) UpdateAvatar(userId int, filepath string) error {
	statement := "UPDATE users SET avatar = ? WHERE id = ?"
	_, err := u.db.Exec(statement, filepath, userId)
	return err
}
