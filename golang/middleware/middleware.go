package middleware

import (
	"net/http"
	"root/gleam/golang/tool/auth"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
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

func CorsProtection() gin.HandlerFunc {
    return func(c *gin.Context) {
        // c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
        c.Next()
    }
}

func CsrfProtection() gin.HandlerFunc {
    return func(c *gin.Context) {
        
		csrfToken := c.Request.Header.Get("X-CSRF-TOKEN")
		
		if c.Request.Method != "GET" {
			session := sessions.Default(c)
            t := session.Get("csrf")
			if csrfToken != t {
				c.JSON(http.StatusUnauthorized, "Unauthorized")
				return
			}
        }
        c.Next()
    }
}