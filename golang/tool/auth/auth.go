package auth

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

type AuthInterface interface {
	CreateAuth(int64, *TokenDetails) error
	FetchAuth(string) (int64, error)
	DeleteRefresh(string) error
	DeleteTokens(*AccessDetails) error
}

var _ AuthInterface = &service{}

var ctx = context.Background()

type service struct {
	client *redis.Client
}

func NewAuth(client *redis.Client) *service {
	return &service{client: client}
}

type AccessDetails struct {
	TokenUuid string
	UserId    int64
}

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	TokenUuid    string
	RefreshUuid  string
	AtExpires    int64
	ReExpires    int64
}

// Save token metadata to Redis
func (tk *service) CreateAuth(userId int64, td *TokenDetails) error {
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC
	rt := time.Unix(td.ReExpires, 0)
	now := time.Now()

	atCreated, err := tk.client.Set(ctx, td.TokenUuid, strconv.Itoa(int(userId)), at.Sub(now)).Result()
	if err != nil {
		return err
	}
	rtCreated, err := tk.client.Set(ctx, td.RefreshUuid, strconv.Itoa(int(userId)), rt.Sub(now)).Result()
	if err != nil {
		return err
	}
	if atCreated == "0" || rtCreated == "0" {
		return errors.New("no record inserted")
	}
	return nil
}

// Check the metadata saved
func (tk *service) FetchAuth(tokenUuid string) (int64, error) {
	userid, err := tk.client.Get(ctx, tokenUuid).Result()
	if err != nil {
		return 0, err
	}
	u, _ := strconv.ParseInt(userid, 10, 64)
	return u, nil
}

func (tk *service) DeleteTokens(authD *AccessDetails) error {
	uid := strconv.FormatInt(authD.UserId, 10)
	// get the refresh uuid
	refreshUuid := fmt.Sprintf("%s++%s", authD.TokenUuid, uid)
	// delete access token
	deleteAt, err := tk.client.Del(ctx, authD.TokenUuid).Result()
	if err != nil {
		return err
	}
	//delete refresh token
	deleteRt, err := tk.client.Del(ctx, refreshUuid).Result()
	if err != nil {
		return err
	}
	if deleteAt != 1 || deleteRt != 1 {
		return errors.New("something went wrong")
	}
	return nil
}

func (tk *service) DeleteRefresh(refreshUuid string) error {
	// delete refresh token
	deleted, err := tk.client.Del(ctx, refreshUuid).Result()
	if err != nil || deleted == 0 {
		return err
	}
	return nil
}
