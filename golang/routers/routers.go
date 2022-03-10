package routers

import (
	
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"root/gleam/golang/routers/api/user"
	"root/gleam/golang/tool/auth"
	"root/gleam/golang/middleware"
	"root/gleam/golang/db"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

func Setup(rc *redis.Client, db *dbDriver.DB) *gin.Engine {

	store := cookie.NewStore([]byte("secret"))

	// initialize gin
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
  	r.Use(sessions.Sessions("csrf", store))
	r.Use(middleware.CorsProtection())
	r.Use(middleware.CsrfProtection())
	
	// serves static files
	//r.StaticFS("/route". "../path/to/files")
	
	//setup JWT methods
	var rd = auth.NewAuth(rc)
	var tk = auth.NewToken()
	
	//Setup Handlers
	var service = user.NewProfile(rd, tk, db)

	//test
	r.POST("/try", service.Try)
	
	//CSRF route
	r.GET("/csrf", service.CSRF)


	
	r.POST("/signup", service.Signup)
	r.POST("/login", service.Login)
	r.POST("/refresh", service.Refresh)
	r.POST("/logout", middleware.TokenAuth(), service.Logout)

	u := r.Group("user")
	u.Use(middleware.TokenAuth())
	{
		
		u.POST("/todo", service.CreateTodo)
	}

	return r
}


