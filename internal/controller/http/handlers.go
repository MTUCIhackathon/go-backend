package http

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/MTUCIhackathon/go-backend/internal/controller/http/model"
	"github.com/MTUCIhackathon/go-backend/internal/model/dto"
)

func (ctrl *Controller) Ping(e echo.Context) error {
	return e.String(http.StatusOK, "pong")
}

func (ctrl *Controller) GetTestByName(e echo.Context) error {
	panic("not implemented")
}
func (ctrl *Controller) GetManyTest(e echo.Context) error {
	panic("not implemented")
}

func (ctrl *Controller) CreateResolved(e echo.Context) error {
	panic("not implemented")
}

func (ctrl *Controller) GetResolvedByID(e echo.Context) error {
	panic("not implemented")
}
func (ctrl *Controller) GetManyResolved(e echo.Context) error {
	panic("not implemented")
}

func (ctrl *Controller) GetOldResolvedByID(e echo.Context) error {
	panic("not implemented")
}

func (ctrl *Controller) CreateConsumer(e echo.Context) error {
	var (
		req  model.CreateConsumerRequest
		err  error
		DTO  dto.CreateConsumer
		resp model.CreateConsumerResponse
	)

	err = e.Bind(&req)
	if err != nil {
		ctrl.log.Error("failed to bind request")
		return e.NoContent(http.StatusBadRequest)
	}

	DTO = dto.CreateConsumer{
		Login:    req.Login,
		Password: req.Password,
	}

	data, err := ctrl.srv.CreateConsumer(e, DTO)
	if err != nil {
		ctrl.log.Error("failed to create consumer")
		return e.NoContent(http.StatusInternalServerError)
	}

	resp = model.CreateConsumerResponse{
		AccessToken:  data.AccessToken,
		RefreshToken: data.RefreshToken,
	}

	return e.JSON(http.StatusCreated, resp)

}
func (ctrl *Controller) GetConsumer(e echo.Context) error {
	data, err := ctrl.srv.GetConsumerByID(e)
	if err != nil {
		ctrl.log.Error("failed to get consumer by id")
		return e.NoContent(http.StatusInternalServerError)
	}

	resp := model.GetConsumerResponse{
		ID:        data.ID,
		Login:     data.Login,
		CreatedAt: data.CreatedAt,
	}
	return e.JSON(http.StatusOK, resp)
}
func (ctrl *Controller) UpdateConsumerPassword(e echo.Context) error {
	var (
		req model.UpdatePasswordRequest
		err error
		DTO dto.UpdatePassword
	)

	err = e.Bind(&req)
	if err != nil {
		ctrl.log.Error("failed to bind request")
		return e.NoContent(http.StatusBadRequest)
	}

	DTO = dto.UpdatePassword{
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	}

	err = ctrl.srv.UpdateConsumerPassword(e, DTO)
	if err != nil {
		ctrl.log.Error("failed to update consumer")
		return e.NoContent(http.StatusInternalServerError)
	}
	return e.NoContent(http.StatusNoContent)
}

func (ctrl *Controller) DeleteConsumer(e echo.Context) error {
	err := ctrl.srv.DeleteConsumerByID(e)
	if err != nil {
		ctrl.log.Error("failed to delete consumer")
		return e.NoContent(http.StatusInternalServerError)
	}
	return e.NoContent(http.StatusNoContent)
}

func (ctrl *Controller) Login(e echo.Context) error {
	var (
		req model.LoginRequest
		err error
	)

	if err = e.Bind(&req); err != nil {
		ctrl.log.Error("failed to bind request")
		return e.NoContent(http.StatusBadRequest)
	}

	DTO := dto.Login{
		Login:    req.Login,
		Password: req.Password,
	}

	data, err := ctrl.srv.Login(e, DTO)
	if err != nil {
		ctrl.log.Error("failed to login")
		return e.NoContent(http.StatusInternalServerError)
	}

	resp := model.LoginResponse{
		AccessToken:  data.AccessToken,
		RefreshToken: data.RefreshToken,
	}

	return e.JSON(http.StatusOK, resp)
}

func (ctrl *Controller) RefreshToken(e echo.Context) error {
	DTO, err := ctrl.srv.RefreshToken(e)
	if err != nil {
		ctrl.log.Error("failed to refresh token")
		return e.NoContent(http.StatusInternalServerError)
	}

	resp := model.RefreshTokenResponse{
		AccessToken:  DTO.AccessToken,
		RefreshToken: DTO.RefreshToken,
	}
	return e.JSON(http.StatusOK, resp)
}
func (ctrl *Controller) SendResultOnEmail(e echo.Context) error {
	panic("not implemented")
}
