package model

type (
	CreateConsumerRequest struct {
		Login    string `json:"login"`
		Password string `json:"password"`
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

type (
	UpdatePasswordRequest struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}
)

type (
	LoginRequest struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
)
