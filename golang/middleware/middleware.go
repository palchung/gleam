package middleware

import (
	"net/http"
	"root/gleam/golang/tool/auth"

	"github.com/gin-gonic/gin"
)

//check if the token is still valid
func TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := auth.TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}
