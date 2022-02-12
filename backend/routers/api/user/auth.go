package user

import (
	"fmt"
	"net/http"
	"strconv"
	"thefreepress/model"
	"thefreepress/tool/setting"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
)

// for testing
var user = model.User{
	ID:       1,
	Username: "username",
	Password: "password",
}

func (p *profileHandler) Try(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello world"})
}

func (p *profileHandler) Login(c *gin.Context) {
	var u model.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	// compare user from database
	if user.Username != u.Username || user.Password != u.Password {
		c.JSON(http.StatusUnauthorized, "Wrong login details")
		return
	}
	ts, err := p.tk.CreateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	saveErr := p.rd.CreateAuth(user.ID, ts)
	if saveErr != nil {
		c.JSON(http.StatusUnprocessableEntity, saveErr.Error())
		return
	}
	tokens := map[string]string{
		"access_token":  ts.AccessToken,
		"refresh_token": ts.RefreshToken,
	}
	c.JSON(http.StatusOK, tokens)
}

func (p *profileHandler) Logout(c *gin.Context) {
	//if metadata is passed and the token valid, delete them from the redis store
	metadata, _ := p.tk.ExtractTokenMetadata(c.Request)
	if metadata != nil {
		deleteErr := p.rd.DeleteTokens(metadata)
		if deleteErr != nil {
			c.JSON(http.StatusBadRequest, deleteErr.Error())
		}
	}
	c.JSON(http.StatusOK, "Successfully Logged out")
}

func (p *profileHandler) Refresh(c *gin.Context) {
	mapToken := map[string]string{}
	if err := c.ShouldBindJSON(&mapToken); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	refreshToken := mapToken["refresh_token"]

	//verify the token
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected siging method: %v", token.Header["alg"])
		}
		return []byte(setting.AppSetting.JWTRefreshSecret), nil
	})

	//if there is an error, the token must have expired
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Refresh token expired")
		return
	}
	// is token valid ?
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		c.JSON(http.StatusUnauthorized, err)
		return
	}
	//Since token is valid, get the uuid
	claims, ok := token.Claims.(jwt.MapClaims) // token claims should conform to the MapClaims
	if ok && token.Valid {
		refreshUuid, ok := claims["refresh_uuid"].(string) // convert the interface into string
		if !ok {
			c.JSON(http.StatusUnprocessableEntity, err)
			return
		}
		userId, roleOk := strconv.ParseInt(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if roleOk != nil {
			c.JSON(http.StatusUnprocessableEntity, "unauthorized")
			return
		}
		//Delete the previous Refresh Token
		delErr := p.rd.DeleteRefresh(refreshUuid)
		if delErr != nil { // if any goes wrong
			c.JSON(http.StatusUnauthorized, "unauthorized")
			return
		}
		//Create new pairs of refresh and access token
		ts, createErr := p.tk.CreateToken(userId)
		if createErr != nil {
			c.JSON(http.StatusForbidden, createErr.Error())
			return
		}
		//save tokens metadata to redis
		saveErr := p.rd.CreateAuth(userId, ts)
		if saveErr != nil {
			c.JSON(http.StatusForbidden, saveErr.Error())
			return
		}
		tokens := map[string]string{
			"access_token":  ts.AccessToken,
			"refresh_token": ts.RefreshToken,
		}
		c.JSON(http.StatusCreated, tokens)
	} else {
		c.JSON(http.StatusUnauthorized, "refresh expired")
	}
}
