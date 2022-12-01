package service

import (
	"context"
	"time"

	"github.com/IRFAN374/upSvc2/models"
	service "github.com/IRFAN374/upSvc2/token"
	log "github.com/go-kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next   service.Service
}

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next service.Service) service.Service {
		return &loggingMiddleware{
			logger: logger,
			next:   next,
		}
	}
}

func (M loggingMiddleware) CreateToken(arg0 context.Context, arg1 uint64, arg2 string) (res0 models.TokenResponse, res1 error) {

	defer func(begin time.Time) {
		M.logger.Log(
			"method", "CreateToken",
			"request", logCreateTokenRequest{
				UserId:   arg1,
				UserName: arg2,
			},
			"response", logCreateTokenResponse{
				Token: res0,
			},
			"err", res1,
			"took", time.Since(begin),
		)

	}(time.Now())
	return M.next.CreateToken(arg0, arg1, arg2)
}

func (M loggingMiddleware) VerifyToken(arg0 context.Context, arg1 string) (res0 bool, res1 error) {

	defer func(begin time.Time) {
		M.logger.Log(
			"method", "VerifyToken",
			"request", logVerifyTokenRequest{
				Token: arg1,
			},
			"response", logVerifyTokenResponse{
				Ok: res0,
			},
			"err", res1,
			"took", time.Since(begin),
		)

	}(time.Now())
	return M.next.VerifyToken(arg0, arg1)
}

type (
	logCreateTokenRequest struct {
		UserId   uint64
		UserName string
	}
	logCreateTokenResponse struct {
		Token models.TokenResponse
	}

	logVerifyTokenRequest struct {
		Token string
	}
	logVerifyTokenResponse struct {
		Ok bool
	}
)
