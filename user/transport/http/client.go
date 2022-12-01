package transporthttp

import (
	"net/url"

	transport "github.com/IRFAN374/upSvc2/user/transport"
	httpkit "github.com/go-kit/kit/transport/http"
)

func NewHTTPClient(u *url.URL, opts ...httpkit.ClientOption) transport.EndpointsSet {

	return transport.EndpointsSet{
		RegisterEndpoint: httpkit.NewClient(
			"POST", u,
			Encode_Register_Request,
			Decode_Register_Response,
			opts...,
		).Endpoint(),
		LoginEndpoint: httpkit.NewClient(
			"POST", u,
			Encode_Login_Request,
			Decode_Login_Response,
			opts...,
		).Endpoint(),
	}

}
