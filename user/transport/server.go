package transport

import (
	"context"

	user "github.com/IRFAN374/upSvc2/user"
	endpoint "github.com/go-kit/kit/endpoint"
)

func Endpoints(svc user.Service) EndpointsSet {
	return EndpointsSet{
		RegisterEndpoint: RegisterEndpoint(svc),
		LoginEndpoint:    LoginEndpoint(svc),
	}
}

func RegisterEndpoint(svc user.Service) endpoint.Endpoint {

	return func(arg0 context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*RegisterRequest)

		res0, res1 := svc.Register(arg0, req.UserName, req.Password)
		return &RegisterResponse{
			UserId: res0,
		}, res1
	}

}

func LoginEndpoint(svc user.Service) endpoint.Endpoint {
	return func(arg0 context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*LoginRequest)

		res0, res1 := svc.Login(arg0, req.UserName, req.Password)
		return &LoginResponse{
			UserId: res0,
		}, res1
	}
}
