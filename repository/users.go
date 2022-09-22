package repository

import (
	"errors"
	"github.com/nuriansyah/log-mbkm-unpas/src"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"regexp"
	"strings"

	_ "github.com/lib/pq"
)

type UserRepository struct {
	db *src.Config
}

func NewUserRepository(db *src.Config) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) Login(email, password string) (*int, error) {
	sqlStatement := "SELECT id,email,password FROM users WHERE id = ?"
	res := u.db.DB.QueryRow(sqlStatement, email, password)

	var hashedPassword string
	var id int
	res.Scan(&id, &hashedPassword)
	if bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) != nil {
		return nil, errors.New("failed Password")
	}
	return &id, nil
}
func (u *UserRepository) CheckEmail(email string) (bool, error) {
	sqlStatement := "SELECT count(*) FROM users WHERE email =?"
	res := u.db.DB.QueryRow(sqlStatement, email)
	var count int
	err := res.Scan(&count)
	if count > 0 {
		return false, err
	}
	return true, err
}
func (u *UserRepository) GetUserRole(id int) (*string, error) {
	statement := "SELECT role FROM users WHERE id = ?"
	var role string
	res := u.db.DB.QueryRow(statement, id)
	err := res.Scan(&role)
	return &role, err
}

func (u *UserRepository) InserNewUser(name, email, role, password string) (usersId int, responCode int, err error) {
	if strings.ToLower(role) != "mahasiswa" && strings.ToLower(role) != "dosen" {
		return -1, http.StatusBadRequest, errors.New("role must to be mahasiswa or dosen")
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

	sqlStetament := "INSERT INTO users(name,email,role,password) VALUES(?,?,?,?)"
	res, err := u.db.DB.Exec(sqlStetament, name, email, strings.ToLower(role), hashedPassword)
	if err != nil {
		return -1, http.StatusBadRequest, err
	}
	resId, err := res.LastInsertId()
	if err != nil {
		return -1, http.StatusBadRequest, err
	}
	return int(resId), http.StatusCreated, err
}
func (u *UserRepository) GetUserData(id int) (*User, error) {
	statement := "SELECT users.id, name, email, role,nrp,prodi, avatar, company, program, batch FROM user_details JOIN users ON users.id = user_details.user_id WHERE users.id = ?"
	var user User
	res := u.db.DB.QueryRow(statement, id)
	err := res.Scan(&user.Id, &user.Name, &user.Email, &user.Role, &user.Nrp, &user.Prodi, &user.Avatar, &user.Company, &user.Program, &user.Batch)
	return &user, err
}
func (u *UserRepository) UpdateDetailDataUser(userID, batch int, nrp, prodi, program, company string) error {
	sqlStmt := `UPDATE user_details SET nrp = ?,prodi = ?,program = ?,company = ?,batch = ? WHERE user_id = ?`
	tx, err := u.db.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(sqlStmt, nrp, prodi, program, company, batch, userID)
	if err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
