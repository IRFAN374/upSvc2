package transport

type (
	LoginRequest struct {
		UserName string `json:"userName"`
		Password string `json:"password"`
	}
	LoginResponse struct {
		UserId string `json:"userId"`
	}

	RegisterRequest struct {
		UserName string `json:"userName"`
		Password string `json:"password"`
	}
	RegisterResponse struct {
		UserId string `json:"userId"`
	}
)
