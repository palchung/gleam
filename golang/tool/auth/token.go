package auth

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/twinj/uuid"
	"net/http"
	"strings"
	"time"
	"root/gleam/golang/tool/setting"
	"strconv"
)

type tokenservice struct {}

func NewToken() *tokenservice {
	return &tokenservice{}
}

type TokenInterface interface {
	CreateToken(userId int64) (*TokenDetails, error)
	ExtractTokenMetadata(*http.Request) (*AccessDetails, error)
}

//token implement the TokenInterface
var _ TokenInterface = &tokenservice{}

func (t *tokenservice) CreateToken(userId int64) (*TokenDetails, error) {
	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 30).Unix() //expires after 30 mins
	td.TokenUuid = uuid.NewV4().String()

	td.ReExpires = time.Now().Add(time.Hour * 24 * 7).Unix() // expire after 1 week
	td.RefreshUuid = uuid.NewV4().String() + "++" + strconv.Itoa(int(userId))
	
	var err error 
	//Create Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.TokenUuid
	atClaims["user_id"] = userId
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(setting.AppSetting.JWTAccessSecret))
	if err != nil {
		return nil, err
	}

	//Create Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_id"] = userId
	rtClaims["exp"] = td.ReExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(setting.AppSetting.JWTRefreshSecret))
	if err != nil {
		return nil, err
	}	
	return td, nil
}

func TokenValid(r *http.Request) error {
	token, err := verifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func verifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token)(interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected siging method: %v", token.Header["alg"])	
		}
		return []byte(setting.AppSetting.JWTAccessSecret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

//get the token from request body
func extractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len (strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func extract(token *jwt.Token) (*AccessDetails, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		userId, userOk := strconv.ParseInt(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if ok == false || userOk != nil {
			return nil, errors.New("unauthorized")
		} else {
			return &AccessDetails{
				TokenUuid: 	accessUuid,
				UserId: 	userId,
			}, nil
		}
	}
	return nil, errors.New("something went wrong")
}

func (t *tokenservice) ExtractTokenMetadata(r *http.Request) (*AccessDetails, error) {
	token, err := verifyToken(r)
	if err != nil {
		return nil, err
	}
	acc, err := extract(token)
	if err != nil {
		return nil, err
	}
	return acc, nil
}
