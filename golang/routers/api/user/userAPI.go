package user

import (
	"root/gleam/golang/db"
	"root/gleam/golang/repository"
	"root/gleam/golang/repository/dbrepo"
	"root/gleam/golang/tool/auth"

	"github.com/gin-gonic/gin"
	"net/http"
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

func checkUnauthorized(err error, c *gin.Context) {
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	return
}

func GetUserIdByToken(p *profileHandler, c *gin.Context) int64 {

	metadata, err := p.tk.ExtractTokenMetadata(c.Request)
	checkUnauthorized(err, c)

	userId, err := p.rd.FetchAuth(metadata.TokenUuid)
	checkUnauthorized(err, c)

	return userId
}
