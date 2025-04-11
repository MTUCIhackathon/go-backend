package model

type (
	CreateUserRequest struct {
		Email string `json:"email"`
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
