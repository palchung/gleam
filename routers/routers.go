package routers

import (
	
	"github.com/gin-gonic/gin"
	"thefreepress/routers/api/user"
	
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	//r.StaticFS("/route". "../path/to/files")
	r.POST("/login", user.Login)

	u := r.Group("user")
	{
		u.POST("/todo", user.CreateTodo)
	}
	return r
}


