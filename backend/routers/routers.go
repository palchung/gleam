package routers

import (
	
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	
	"thefreepress/routers/api/user"
	"thefreepress/tool/auth"
	"thefreepress/middleware"
)

func InitRouter(client *redis.Client) *gin.Engine {
	
	var rd = auth.NewAuth(client)
	var tk = auth.NewToken()
	var service = user.NewProfile(rd, tk)

	// initialize gin
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// serves static files
	//r.StaticFS("/route". "../path/to/files")

	// serves public api
	r.POST("/login", service.Login)

	u := r.Group("user")
	u.Use(middleware.TokenAuth())
	{
		u.POST("/todo", service.CreateTodo)
	}

	return r
}


