package transporthttp

import (
	http1 "net/http"

	transport "github.com/IRFAN374/upSvc2/user/transport"
	http "github.com/go-kit/kit/transport/http"
	mux "github.com/gorilla/mux"
)

func NewHTTPHandler(endpoints *transport.EndpointsSet, opts ...http.ServerOption) http1.Handler {
	mux := mux.NewRouter()

	mux.Methods("POST").Path("/api/register").Handler(
		http.NewServer(
			endpoints.RegisterEndpoint,
			Decode_Register_Request,
			Encode_Register_Response,
			opts...))
	mux.Methods("POST").Path("/api/login").Handler(
		http.NewServer(
			endpoints.LoginEndpoint,
			Decode_Login_Request,
			Encode_Login_Response,
			opts...))

	return mux

}
