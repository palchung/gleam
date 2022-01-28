package user

import (
	"github.com/gin-gonic/gin"
)

type Todo struct {
	UserID 	uint64 	`json:"user_id"`
	Title 	string	`json:"title"`	
}


func CreateTodo(c *gin.Context) {
	
}