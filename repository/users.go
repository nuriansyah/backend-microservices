package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"regexp"
	"strings"

	_ "github.com/lib/pq"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

/*
ouh i just noticed, this not how to use bcrypt.CompareHashAndPassword
it return an error, you need to check if its type of bcrypt.ErrMismatchedHashAndPassword
func main() {
    p, _ := bcrypt.GenerateFromPassword([]byte("abc"), bcrypt.DefaultCost)
    err := bcrypt.CompareHashAndPassword(p, []byte("abc"))
    if err != nil {
        // check if error is type of password missmatch
        if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
            fmt.Println("wrong password")
            return
        }
        // panic on other error
        panic(err)
    }
    fmt.Println("wellcome")
}
*/

func (u *UserRepository) Login(email, password string) (*int, error) {
	sqlStatement := "SELECT id, password FROM users WHERE email = $1"
	res := u.db.QueryRow(sqlStatement, email, password)
	var hashedPassword string
	var id int
	res.Scan(&id, &hashedPassword)
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			fmt.Println("Wrong Passowrd")
			return nil, err
		}
		panic(err)
	}
	return &id, nil
}
func (u *UserRepository) CheckEmail(email string) (bool, error) {
	sqlStatement := "SELECT count(*) FROM users WHERE email = $1"
	res := u.db.QueryRow(sqlStatement, email)
	var count int
	err := res.Scan(&count)
	if count > 0 {
		return false, err
	}
	return true, err
}
func (u *UserRepository) GetUserRole(id int) (*string, error) {
	statement := "SELECT role FROM users WHERE id = $1"
	var role string
	res := u.db.QueryRow(statement, id)
	err := res.Scan(&role)
	return &role, err
}

func (u *UserRepository) GetUserData(id int) (*User, error) {
	statement := `SELECT users.id, name, email, role,nrp,prodi, avatar, company, program, batch FROM user_details JOIN users ON users.id = user_details.user_id WHERE users.id = $1`
	var user User
	res := u.db.QueryRow(statement, id)
	err := res.Scan(&user.Id, &user.Name, &user.Email, &user.Role, &user.Nrp, &user.Prodi, &user.Avatar, &user.Company, &user.Program, &user.Batch)
	return &user, err
}
func (u *UserRepository) UpdateDetailDataUser(userID, batch int, nrp, prodi, program, company string) error {
	sqlStmt := `UPDATE user_details SET nrp = $1,prodi = $2,program = $3,company = $4,batch = $5 WHERE user_id = $6`
	tx, err := u.db.Begin()
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

func (u *UserRepository) InsertUser(name, email, password, role string) (userId, responCode int, err error) {
	if strings.ToLower(role) != "mahasiswa" && strings.ToLower(role) != "siswa" {
		return -1, http.StatusBadRequest, errors.New("role must be either 'mahasiswa' or 'siswa'")
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

	s := string([]byte(hashedPassword))
	println(s)

	println(&hashedPassword)
	sqlStatement := `INSERT INTO users (name,email,password,role) VALUES ($1,$2,$3,$4) RETURNING id`

	var id int
	err = u.db.QueryRow(sqlStatement, name, email, string(hashedPassword), strings.ToLower(role)).Scan(&id)

	//stmt, err := u.db.Prepare(sqlStatement)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer stmt.Close()
	//
	//var id int
	//err = stmt.QueryRow(name, email, hashedPassword, role).Scan(&id)
	//if err != nil {
	//	log.Fatal(err)
	//}

	return userId, http.StatusOK, err
}
