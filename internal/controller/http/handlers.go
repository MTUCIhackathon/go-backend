package http

import (
	"net/http"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/labstack/echo/v4"

	"github.com/MTUCIhackathon/go-backend/internal/controller/http/model"
	"github.com/MTUCIhackathon/go-backend/internal/model/dto"
)

func (ctrl *Controller) Ping(e echo.Context) error {
	return e.String(http.StatusOK, "pong")
}

func (ctrl *Controller) GetTestByID(e echo.Context) error {
	var (
		err      error
		response model.GetTestResponse
	)
	testID := e.Param("test_id")
	token := e.Request().Header.Get(echo.HeaderAuthorization)
	id, err := uuid.Parse(testID)
	if err != nil {
		ctrl.log.Error("failed to parse test id", zap.Error(err))
	}
	ctrl.log.Debug("get test_id from path", zap.Any("id", testID))

	resp, err := ctrl.srv.GetTestByID(e.Request().Context(), token, id)
	if err != nil {
		ctrl.log.Error("failed to get by test_id", zap.Any("id", id), zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get test by id")
	}

	questions := make([]model.TestQuestion, len(resp.Questions))
	for i := 0; i < len(resp.Questions); i++ {
		questions[i] = model.TestQuestion{
			Order:    resp.Questions[i].Order,
			Question: resp.Questions[i].Question,
		}
	}

	response = model.GetTestResponse{
		ID:        resp.ID,
		Name:      resp.Name,
		Questions: questions,
	}

	return e.JSON(http.StatusOK, response)
}
func (ctrl *Controller) GetAllTest(e echo.Context) error {
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

	data, err := ctrl.srv.CreateConsumer(e.Request().Context(), DTO)
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
	token := e.Request().Header.Get(echo.HeaderAuthorization)
	data, err := ctrl.srv.GetConsumerByID(e.Request().Context(), token)
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

	token := e.Request().Header.Get(echo.HeaderAuthorization)

	DTO = dto.UpdatePassword{
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
		Token:       token,
	}

	err = ctrl.srv.UpdateConsumerPassword(e.Request().Context(), DTO)
	if err != nil {
		ctrl.log.Error("failed to update consumer")
		return e.NoContent(http.StatusInternalServerError)
	}
	return e.NoContent(http.StatusNoContent)
}

func (ctrl *Controller) DeleteConsumer(e echo.Context) error {
	token := e.Request().Header.Get(echo.HeaderAuthorization)

	err := ctrl.srv.DeleteConsumerByID(e.Request().Context(), token)
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

	data, err := ctrl.srv.Login(e.Request().Context(), DTO)
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
	token := e.Request().Header.Get(echo.HeaderAuthorization)

	DTO, err := ctrl.srv.RefreshToken(e.Request().Context(), token)
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
