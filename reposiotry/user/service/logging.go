package service

import (
	"context"
	"time"

	"github.com/IRFAN374/upSvc2/models"
	service "github.com/IRFAN374/upSvc2/reposiotry/user"
	log "github.com/go-kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next   service.Repository
}

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next service.Repository) service.Repository {
		return &loggingMiddleware{
			logger: logger,
			next:   next,
		}
	}
}

func (M loggingMiddleware) Add(arg0 context.Context, arg1 string, arg2 models.User) (res0 error) {

	defer func(begin time.Time) {
		M.logger.Log(
			"method", "Add",
			"request", logAddRequest{
				UserId:   arg1,
				UserInfo: arg2,
			},
			"err", res0,
			"took", time.Since(begin),
		)

	}(time.Now())
	return M.next.Add(arg0, arg1, arg2)

}

func (M loggingMiddleware) Get(arg0 context.Context, arg1 string) (res0 models.User, res1 error) {

	defer func(begin time.Time) {
		M.logger.Log(
			"method", "Get",
			"request", logGetRequest{
				UserId: arg1,
			},
			"response", logGetResponse{
				UserInfo: res0,
			},
			"err", res1,
			"took", time.Since(begin),
		)

	}(time.Now())
	return M.next.Get(arg0, arg1)
}

func (M loggingMiddleware) Update(arg0 context.Context, arg1 string) (res0 models.User, res1 error) {

	defer func(begin time.Time) {
		M.logger.Log(
			"method", "Update",
			"request", logUpdateRequest{
				UserId: arg1,
			},
			"response", logUpdateResponse{
				UserInfo: res0,
			},
			"err", res1,
			"took", time.Since(begin),
		)

	}(time.Now())
	return M.next.Update(arg0, arg1)
}

func (M loggingMiddleware) Delete(arg0 context.Context, arg1 string) (res0 error) {

	defer func(begin time.Time) {
		M.logger.Log(
			"method", "Delete",
			"request", logDeleteRequest{
				UserId: arg1,
			},
			"err", res0,
			"took", time.Since(begin),
		)

	}(time.Now())
	return M.next.Delete(arg0, arg1)

}

func (M loggingMiddleware) IsExist(arg0 context.Context, arg1 string) (res0 bool, res1 error) {

	defer func(begin time.Time) {
		M.logger.Log(
			"method", "IsExist",
			"request", logIsExistRequest{
				UserId: arg1,
			},
			"response", logIsExistResponse{
				Ok: res0,
			},
			"err", res1,
			"took", time.Since(begin),
		)

	}(time.Now())
	return M.next.IsExist(arg0, arg1)
}

type (
	logAddRequest struct {
		UserId   string
		UserInfo models.User
	}

	logGetRequest struct {
		UserId string
	}
	logGetResponse struct {
		UserInfo models.User
	}

	logUpdateRequest struct {
		UserId string
	}
	logUpdateResponse struct {
		UserInfo models.User
	}

	logDeleteRequest struct {
		UserId string
	}

	logIsExistRequest struct {
		UserId string
	}
	logIsExistResponse struct {
		Ok bool
	}
)
