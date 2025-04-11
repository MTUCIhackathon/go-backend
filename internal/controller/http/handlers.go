package http

import (
	"github.com/MTUCIhackathon/go-backend/internal/controller/http/model"
	"github.com/MTUCIhackathon/go-backend/internal/model/dto"
	"github.com/labstack/echo/v4"
	"net/http"
)

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
		Email:    req.Email,
		Password: req.Password,
	}

	data, err := ctrl.srv.Consumer().CreateConsumer(e, DTO)
	if err != nil {
		ctrl.log.Error("failed to create consumer")
		return e.NoContent(http.StatusInternalServerError)
	}

	resp = model.CreateConsumerResponse{
		AccessToken:  data.AccessToken,
		RefreshToken: data.RefreshToken,
	}

	return e.JSON(http.StatusOK, resp)

}
func (ctrl *Controller) GetConsumerByID(e echo.Context) error {
	panic("not implemented")
}
func (ctrl *Controller) UpdateConsumerPassword(e echo.Context) error {
	panic("not implemented")
}
func (ctrl *Controller) SendResultOnEmail(e echo.Context) error {
	panic("not implemented")
}
