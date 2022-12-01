package transport

import "context"

func (set EndpointsSet) Register(arg0 context.Context, arg1 string, arg2 string) (res0 string, res1 error) {

	request := RegisterRequest{
		UserName: arg1,
		Password: arg2,
	}
	response, res1 := set.RegisterEndpoint(arg0, &request)
	if res1 != nil {
		return
	}

	return response.(*RegisterResponse).UserId, res1
}

func (set EndpointsSet) Login(arg0 context.Context, arg1 string, arg2 string) (res0 string, res1 error) {
	request := LoginRequest{
		UserName: arg1,
		Password: arg2,
	}
	response, res1 := set.LoginEndpoint(arg0, &request)
	if res1 != nil {
		return
	}

	return response.(*LoginResponse).UserId, res1
}
