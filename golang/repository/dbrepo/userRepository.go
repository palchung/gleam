package dbrepo

import (
	"root/gleam/golang/model"
	"fmt"
	"github.com/jackc/pgconn"
	"errors"
	"os"
	"time"
	"root/gleam/golang/tool/logging"
)

var pgerr error = nil 
var pgErr *pgconn.PgError

func handleErrors(err error) error{
	if errors.As(err, &pgErr) {
	  logging.Fatal(pgErr.Message)
		return pgErr
	}
}


func (m *postgresDBRepo) AllUsers() bool {
	return true
}

func (m *postgresDBRepo) GetUserPwdByEmail(email string) (userID string, password string, error) {

	var userID, password string
	sqlSmt := `SELECT id, password FROM users WHERE email=$1`
	err := m.DB.QueryRow(sqlSmt, email).Scan(&userID, &password)
	if err != nil {
		pgerr = handleErrors(err)
	}
	return userID, password, pgerr
}

func (m *postgresDBRepo) SaveNewUser(u model.User, hashedPwd string) (userID int64, error) {

	var userID int64
	t := time.Now()
	sqlSmt := `INSERT INTO users(firstName, lastName, email, password, createdAt, updatedAt) VALUES($1, $2, $3, $4, $5, $6) RETURNING id`
	err := m.DB.QueryRow(sqlSmt, u.FirstName, u.LastName, u.Email, hashedPwd, t, t).Scan(&userID)	
	if err != nil {
		pgerr = handleErrors(err)
	}
	return userID, pgerr
}
