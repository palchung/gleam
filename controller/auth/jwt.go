package auth

import (
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/twinj/uuid"
	"github.com/go-redis/redis/v8"
	"time"
	"os"
)

var redisClient = *redis.Client

type TokenDetails struct {
	AccessToken		string
	RefreshToken	string
	AccessUuid		string
	RefreshUuid		string
	AtExpires		int64
	RtExpires		int64
}

type AccessDetails struct {
	AccessUuid 	string
	UserId 		uint64
}


func CreateJWTToken(userid uint64) (*TokenDetails, error) {
	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(time,Minute *15).Unix()
	td.AccessUuid = uuid.NewV4().String()
	
	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = uuid.NewV4().String()
	
	var err error
	//Create Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["user_id"] = userid
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err := at.SignedString([]byte(os.Getenv("JWTAccessSecret")))
	if err != nil {
		return nil, err
	}
	
	//Create Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_id"] = userid
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("JWTRefreshSecret")))
	if err != nil {
		return nil, err
	}
	return td, nil
}

func CreateJWTAuth(userid uint64, td *TokenDetails) error {
	at := time.Unix(td.AtExpires,0) // convert Unix to UTC
	rt := time.Unix(td.RtExpires,0)
	now := time.Now()

	errAccess := redisClient.Set(td.AccessUuid, strconv.Itoa(int(userid)), at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	errRefresh := redisClient.Set(td.RefreshUuid, strconv.Itoa(int(userid)), rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := kwt.Parse(tokenString, func(token *jwt.Token (interface{}, error) {
		// make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpect signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWTAccessSecret")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func ExtractTokenMetadata(r *http.Request) (*AccessDetails, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		return &AccessDetails {
			AccessUuid: accessUuid,
			UserId: userId,
		}, nil
	}
	return nil , err
	
}

func FetchAuth(authD *AccessDetails) (uint64, error) {
	userid, err := redisClient.Get(authD.AccessUuid).Result()
	if err != nil {
		return 0, err
	}
	userID, _ := strconv.ParseUint(userid, 10, 64)
	return userID, nil
}

func DeleteAuth(givenUuid string) (int64, error) {
	deleted, err := redisClient.Del(givenUuid).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		c.Next()
	}
}


