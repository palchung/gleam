package user

import (
	"fmt"
	"net/http"
	"strconv"
	"root/gleam/golang/model"
	"root/gleam/golang/tool/setting"
	"root/gleam/golang/tool/password"
	"root/gleam/golang/tool/logging"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
)

func (p *profileHandler) Try(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello world"})
}

// context: firstName, lastName, email, password
func (p *profileHandler) Signup(c *gin.Context){
	var u model.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	
	hashedPwd := password.HashAndSalt(u.Password)
	
	// save user into database
	// return user id
	userID,errDB := p.db.SaveNewUser(u, hashedPwd)
	if errDB != nil {
		var errorMsg string
		if errDB.Code == "23505" {
			errorMsg = "User Account Exist"
    	} else {
			errorMsg = "Internal Error"
		}
		c.JSON(http.StatusUnprocessableEntity, errorMsg)
		return
	}
	
	ts, err := p.tk.CreateToken(userID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	saveErr := p.rd.CreateAuth(userID, ts)
	if saveErr != nil {
		c.JSON(http.StatusUnprocessableEntity, saveErr.Error())
		return
	}
	res := map[string]string{
		"access_token":  ts.AccessToken,
		"userID": strconv.FormatInt(userID, 10),
	}	
	c.SetCookie("refresh_token", ts.RefreshToken, 60*60*60, "/", setting.AppSetting.PrefixUrl, true, true)
	c.JSON(http.StatusOK, res)
}



func (p *profileHandler) Login(c *gin.Context) {
	var u model.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	
	userID, userPwd, errDB := p.db.GetUserPwdByEmail(u.Email)
	if errDB != nil {
		c.JSON(http.StatusUnauthorized, "We can't log you in")
		return
	}

	verify := password.Verify(userPwd, u.Password)
	if !verify {
		c.JSON(http.StatusUnauthorized, "We can't log you in")
		return
	}
	
	ts, err := p.tk.CreateToken(userID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	saveErr := p.rd.CreateAuth(userID, ts)
	if saveErr != nil {
		c.JSON(http.StatusUnprocessableEntity, saveErr.Error())
		return
	}

	res := map[string]string{
		"access_token":  ts.AccessToken,
		"userID": strconv.FormatInt(userID, 10),
	}	
	c.SetCookie("refresh_token", ts.RefreshToken, 60*60*60, "/", setting.AppSetting.PrefixUrl, true, true)
	c.JSON(http.StatusOK, res)
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
	// mapToken := map[string]string{}
	// if err := c.ShouldBindJSON(&mapToken); err != nil {
	// 	c.JSON(http.StatusUnprocessableEntity, err.Error())
	// 	return
	// }
	// refreshToken := mapToken["refresh_token"]

	refreshToken, err := c.Cookie("refresh_token")
	err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

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
		res := map[string]string{
			"access_token":  ts.AccessToken,
			"userID": strconv.FormatInt(userID, 10),
		}	
		c.SetCookie("refresh_token", ts.RefreshToken, 60*60*60, "/", setting.AppSetting.PrefixUrl, true, true)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusUnauthorized, "refresh expired")
	}
}
