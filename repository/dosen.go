package repository

import (
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"regexp"
)

type DosenRepository struct {
	db *sql.DB
}

func NewDosenRepository(db *sql.DB) *DosenRepository {
	return &DosenRepository{db: db}
}

func (u *DosenRepository) Login(email string, password string) (*int, error) {
	statement := "SELECT id,email,password  FROM dosen WHERE email = ?"
	res := u.db.QueryRow(statement, email, password)
	var hashedPassword string
	var id int
	res.Scan(&id, &hashedPassword)
	if bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) != nil {
		return nil, errors.New("Login Failed")
	}

	_, err := u.db.Exec(statement, id)
	if err != nil {
		return nil, err
	}
	return &id, nil
}
func (u *DosenRepository) GetUserData(id int) (*Mahasiswa, error) {
	statement := "SELECT id,name,email,avatar FROM mahasiswa "
	var mhs Mahasiswa
	res := u.db.QueryRow(statement, id)
	err := res.Scan(&mhs.Id, &mhs.Name, &mhs.Email, &mhs.Avatar)
	return &mhs, err
}
func (u *DosenRepository) CheckEmail(email string) (bool, error) {
	statement := "SELECT count(*) FROM mahasiswa WHERE email = ?"
	res := u.db.QueryRow(statement, email)

	var count int
	err := res.Scan(&count)
	if count > 0 {
		return false, err
	}
	return true, err
}
func (u *DosenRepository) UpdateAvatar(userId int, filepath string) error {
	statement := "UPDATE dosen SET avatar = ? WHERE id = ?"
	_, err := u.db.Exec(statement, filepath, userId)
	return err
}

func (u *DosenRepository) UpdateDataDosen(id int, name string) error {
	statement := "UPDATE dosen SET name = ? WHERE id = ?"
	_, err := u.GetUserData(id)
	if err != nil {
		return err
	}
	_, err = u.db.Exec(statement, name, id)
	return err
}
func (u *DosenRepository) Register(name, email, password string) (userId int, responseCode int, err error) {
	isAvailble, err := u.CheckEmail(email)
	if err != nil {
		return -1, http.StatusBadRequest, err
	}
	if !isAvailble {
		return -1, http.StatusBadRequest, errors.New("Email has been used")
	}
	regex, err := regexp.Compile("^[\\w-\\.]+@([\\w-]+\\.)+[\\w-]{2,4}$")
	if err != nil {
		return -1, http.StatusInternalServerError, err
	}
	isValid := regex.Match([]byte(email))
	if !isValid {
		return -1, http.StatusBadRequest, errors.New("invalid email")
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	sqlStmt := "INSERT dosen INTO (name,email,password) VALUES (?,?,?)"
	res, err := u.db.Exec(sqlStmt, name, hashedPassword)
	if err != nil {
		return -1, http.StatusInternalServerError, err
	}
	if err != nil {
		return -1, http.StatusInternalServerError, err
	}
	resId, err := res.LastInsertId()
	if err != nil {
		return -1, http.StatusInternalServerError, err
	}
	return int(resId), http.StatusOK, err
}
