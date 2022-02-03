package user

import (
	"thefreepress/tool/auth"
	"thefreepress/repository"
	"thefreepress/repository/dbrepo"
	"thefreepress/db"
)


type profileHandler struct {
	rd	auth.AuthInterface
	tk	auth.TokenInterface
	db	repository.DatabaseInterface
}

func NewProfile(rd auth.AuthInterface, tk auth.TokenInterface, db *dbDriver.DB) *profileHandler {
	return &profileHandler{
		rd, 
		tk, 
		dbrepo.NewPostgresRepo(db.SQL),
	}
}
