package model

type (
	CreateConsumerRequest struct {
		Login    string  `json:"login"`
		Email    *string `json:"email"`
		Password string  `json:"password"`
	}

	UpdatePasswordRequest struct {
		Token string `json:"token"`
	}
)

type (
	GetTestRequest struct {
		Name string `json:"name"`
	}
	GetManyTestRequest struct {
		// TODO ?
		Names []string `json:"names"`
	}
)
