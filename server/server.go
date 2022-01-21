package server

import(
	"github.com/gin-gonic/gin"
)


type IController interface {
	Router(server *Server)
}

type Server struct {
	*gin.Engine
	g *gin.RouterGroup
}

func Init() *Server {
	s := &Server{Engine: gin.New()}
	return s
}

func (this *Server) Listen(port string){
	this.Run(":" + port)
}



func (this *Server) Route(controllers ...IController) *Server {
	for _, c := range controllers {
		c.Router(this)
	}
	return this
}

func (this *Server) UserRoute(group string, controllers ...IController) *Server {
	this.g = this.Group(group)
	for _, c := range controllers {
		c.Router(this)
	}
	return this
}