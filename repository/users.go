package repository

import (
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"regexp"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) Login(email string, password string) (*int, error) {
	statement := "SELECT id, password FROM users WHERE email = ?"
	res := u.db.QueryRow(statement, email, password)
	var hashedPassword string
	var id int
	res.Scan(&id, &hashedPassword)
	if bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) != nil {
		return nil, errors.New("Login Failed")
	}
	return &id, nil
}

func (u *UserRepository) GetUserRole(id int) (*string, error) {
	statement := "SELECT role FROM users WHERE id = ?"
	var role string
	res := u.db.QueryRow(statement, id)
	err := res.Scan(&role)
	return &role, err
}

func (u *UserRepository) CheckEmail(email string) (bool, error) {
	sqlStatement := "SELECT count(*) FROM users WHERE email = ?"
	res := u.db.QueryRow(sqlStatement, email)
	var count int
	err := res.Scan(&count)
	if count > 0 {
		return false, err
	}
	return true, err
}
func (u *UserRepository) InsertNewUser(name string, email string, password string, role string, program *string, company *string, batch *int) (userId, responseCode int, err error) {
	if strings.ToLower(role) != "mahasiswa" && strings.ToLower(role) != "dosen" {
		return -1, http.StatusBadRequest, errors.New("role must be either 'mahasiswa' or 'dosen'")
	}
	if strings.ToLower(role) == "mahasiswa" {
		if program == nil || batch == nil {
			return -1, http.StatusBadRequest, errors.New("please fill program and batch correctly")
		}
	}
	isAvailable, err := u.CheckEmail(email)
	if err != nil {
		return -1, http.StatusBadRequest, err
	}

	if !isAvailable {
		return -1, http.StatusBadRequest, errors.New("email has been used")
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
	sqlStatement := "INSERT INTO users (name, email, password, role) VALUES (?, ?, ?, ?)"
	res, err := u.db.Exec(sqlStatement, name, email, hashedPassword, strings.ToLower(role))
	if err != nil {
		return -1, http.StatusInternalServerError, err
	}
	resId, err := res.LastInsertId()
	if err != nil {
		return -1, http.StatusInternalServerError, err
	}
	sqlStatement = "INSERT INTO user_details (user_id,program,company,batch) VALUES (?,?,?,?)"
	_, err = u.db.Exec(sqlStatement, resId, program, company, batch)
	return int(resId), http.StatusCreated, err
}
