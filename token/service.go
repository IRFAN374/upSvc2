package token

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/IRFAN374/upSvc2/models"
	"github.com/IRFAN374/upSvc2/reposiotry/token"
	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
)

type Service interface {
	CreateToken(ctx context.Context, userid uint64, username string) (token models.TokenResponse, err error)
	VerifyToken(ctx context.Context, token string) (ok bool, err error)
}

type service struct {
	tokenRepo token.Repository
}

func NewService(tokenRepo token.Repository) *service {
	return &service{
		tokenRepo: tokenRepo,
	}
}

func (svc *service) CreateToken(ctx context.Context, userid uint64, username string) (token models.TokenResponse, err error) {

	td := &models.TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUuid = uuid.NewV4().String()
	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = uuid.NewV4().String()

	atClaims := jwt.MapClaims{}
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd")
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["user_id"] = userid
	atClaims["exp"] = td.AtExpires

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return models.TokenResponse{}, err
	}
	//Creating Refresh Token
	os.Setenv("REFRESH_SECRET", "mcmvmkmsdnfsdmfdsjf") //this should be in an env file
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_id"] = userid
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return models.TokenResponse{}, err
	}

	err =svc.tokenRepo.Add(ctx, strconv.Itoa(int(userid)), td)

	if err != nil {
		return models.TokenResponse{}, err
	}

	return 
}

func (svc *service) VerifyToken(ctx context.Context, token string) (ok bool, err error) {
	return
}
