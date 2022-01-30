package auth

import (
	"errors"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type AuthInterface interface {
	CreateAuth(string, *TokenDetails) error
	FetchAuth(string) (string, error)
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
	TokenUuid	string
	UserId		string
}

type TokenDetails struct {
	AccessToken		string
	RefreshToken	string
	TokenUuid		string
	RefreshUuid		string
	AtExpires		int64
	ReExpires		int64
}

// Save token metadata to Redis
func (tk *service) CreateAuth(userId string, td *TokenDetails) error {
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC
	rt := time.Unix(td.ReExpires, 0)
	now := time.Now()

	atCreated, err := tk.client.Set(ctx, td.TokenUuid, userId, at.Sub(now)).Result()
	if err != nil {
		return err
	}
	rtCreated, err := tk.client.Set(ctx, td.RefreshUuid, userId, rt.Sub(now)).Result()
	if err != nil {
		return err
	}
	if atCreated == "0" || rtCreated == "0" {
		return errors.New("no record inserted")
	}
	return nil
}

// Check the metadata saved
func (tk *service) FetchAuth(tokenUuid string) (string, error) {
	userid, err := tk.client.Get(ctx, tokenUuid).Result()
	if err != nil {
		return "", err
	}
	return userid, nil
}

func (tk *service) DeleteTokens(authD *AccessDetails) error {
	// get the refresh uuid
	refreshUuid := fmt.Sprintf("%s++%s", authD.TokenUuid, authD.UserId)
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