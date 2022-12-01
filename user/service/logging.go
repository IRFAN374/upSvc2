package service

import (
	"context"
	"time"

	service "github.com/IRFAN374/upSvc2/user"
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

func (M loggingMiddleware) Register(arg0 context.Context, arg1 string, arg2 string) (res0 string, res1 error) {

	defer func(begin time.Time) {
		M.logger.Log(
			"method", "Register",
			"request", logRegisterRequest{
				UserName: arg1,
				Password: arg2,
			},
			"response", logRegisterResponse{
				UserId: res0,
			},
			"err", res1,
			"took", time.Since(begin),
		)

	}(time.Now())
	return M.next.Register(arg0, arg1, arg2)
}

func (M loggingMiddleware) Login(arg0 context.Context, arg1 string, arg2 string) (res0 error) {

	defer func(begin time.Time) {
		M.logger.Log(
			"method", "Login",
			"request", logLoginRequest{
				UserName: arg1,
				Password: arg2,
			},
			"response", logLoginResponse{},
			"err", res0,
			"took", time.Since(begin),
		)

	}(time.Now())

	return M.next.Login(arg0, arg1, arg2)
}

type (
	logRegisterRequest struct {
		UserName string
		Password string
	}
	logRegisterResponse struct {
		UserId string
	}

	logLoginRequest struct {
		UserName string
		Password string
	}
	logLoginResponse struct {
	}
)
