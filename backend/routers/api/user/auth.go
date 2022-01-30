package user

import(
	"github.com/gin-gonic/gin"
	"thefreepress/tool/auth"
	
)

type profileHandler struct {
	rd	auth.AuthInterface
	tk	auth.TokenInterface
}

func NewProfile(rd auth.AuthInterface, tk auth.TokenInterface) *profileHandler {
	return &profileHandler{rd, tk}
}




func (p *profileHandler) Login(c *gin.Context) {
	

}

func (p *profileHandler) Logout(c *gin.Context) {
	
}