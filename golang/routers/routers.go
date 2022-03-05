package routers

import (
	
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	
	"root/gleam/golang/routers/api/user"
	"root/gleam/golang/tool/auth"
	"root/gleam/golang/middleware"
	"root/gleam/golang/db"
)

func Setup(rc *redis.Client, db *dbDriver.DB) *gin.Engine {

	// initialize gin
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.CORS())
	
	// serves static files
	//r.StaticFS("/route". "../path/to/files")
	
	//setup JWT methods
	var rd = auth.NewAuth(rc)
	var tk = auth.NewToken()
	
	//Setup Handlers
	var service = user.NewProfile(rd, tk, db)

	//test
	r.POST("/try", service.Try)
	// r.OPTIONS("/try", service.Try)

	// serves public api
	r.POST("/signup", service.Signup)
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


