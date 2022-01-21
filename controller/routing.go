package routing

import (
	"thefreepress/controller/user"
	"thefreepress/controller/auth"
	"thefreepress/server"
)

// type UserController struct {
// }

type UserController struct {
}

// func NewUserController() *UserController {
// 	return &UserController{}
// }

func Authentication() *UserController {
	return &UserController{}
}

// func (this *UserController) GetUser() gin.HandlerFunc {
// 	return func(ctx *gin.Context){
// 		ctx.JSON(200, gin.H{
// 			"data": "hello world",
// 		})
// 	}
// }

// func (this *OrderController) GetOrder() gin.HandlerFunc {
// 	return func(ctx *gin.Context){
// 		ctx.JSON(200, gin.H{
// 			"data": "get order",
// 		})
// 	}
// }


// func (this *UserController) Router (server *server.Server) {
// 	server.Handle("GET", "/user", this.GetUser())
// }

func (this *UserController) Router (server *server.Server) {
	server.Handle("POST", "/login", user.Login)
	server.Handle("POST", "/todo", auth.TokenAuthMiddleware ,user.CreateTodo)
	server.Handle("POST", "logout", auth.TokenAuthMiddleware, user.Logout)
}