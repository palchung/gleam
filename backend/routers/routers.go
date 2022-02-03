package routers

import (
	
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	
	"thefreepress/routers/api/user"
	"thefreepress/tool/auth"
	"thefreepress/middleware"
	"thefreepress/db"
)

func Setup(rc *redis.Client, db *dbDriver.DB) *gin.Engine {

	// initialize gin
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	
	// serves static files
	//r.StaticFS("/route". "../path/to/files")
	
	//setup JWT
	var rd = auth.NewAuth(rc)
	var tk = auth.NewToken()
	
	//Setup Handlers
	var service = user.NewProfile(rd, tk, db)

	//test
	r.GET("/try", service.Try)

	// serves public api
	r.POST("/login", service.Login)
	r.POST("/refresh", service.Refresh)

	u := r.Group("user")
	u.Use(middleware.TokenAuth())
	{
		u.POST("/logout", service.Logout)
		u.POST("/todo", service.CreateTodo)
	}

	return r
}


