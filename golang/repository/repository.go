package repository

import (
	"root/gleam/golang/model"
)

type DatabaseInterface interface {
	AllUsers() bool
	SaveNewUser(model.User, string) (int64, error)
	GetUserPwdByEmail(string) (int64, string, error)
}
