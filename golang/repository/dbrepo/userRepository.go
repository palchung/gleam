package dbrepo

import (
	"root/gleam/golang/model"	
	"time"
)



func (m *postgresDBRepo) AllUsers() bool {
	return true
}

func (m *postgresDBRepo) GetUserPwdByEmail(email string) (userID int64, password string, err error) {

	sqlSmt := `SELECT id, password FROM users WHERE email=$1`
	err = m.DB.QueryRow(sqlSmt, email).Scan(&userID, &password)
	
	return userID, password, err
}

func (m *postgresDBRepo) SaveNewUser(u model.User, hashedPwd string) (userID int64, err error) {

	t := time.Now()
	sqlSmt := `INSERT INTO users(firstName, lastName, email, password, createdAt, updatedAt) VALUES($1, $2, $3, $4, $5, $6) RETURNING id`
	err = m.DB.QueryRow(sqlSmt, u.FirstName, u.LastName, u.Email, hashedPwd, t, t).Scan(&userID)	
	
	return userID, err
}
