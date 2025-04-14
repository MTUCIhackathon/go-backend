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
		response model.GetTestResponse
	)

	testID := e.Param("test_id")
	token := e.Request().Header.Get(echo.HeaderAuthorization)

	id, err := uuid.Parse(testID)
	if err != nil {
		ctrl.log.Error(
			"failed to parse test id",
			zap.Error(err),
		)

		return echo.NewHTTPError(http.StatusBadRequest)
	}

	ctrl.log.Debug("get test id from path", zap.Any("id", testID))

	resp, err := ctrl.srv.GetTestByID(e.Request().Context(), token, id)
	if err != nil {
		ctrl.log.Error(
			"failed to get test by id",
			zap.String("id", id.String()),
			zap.Error(err),
		)

		return handleErr(err)
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

	ctrl.log.Debug("successfully handle get test by id")

	return e.JSON(http.StatusOK, response)
}

func (ctrl *Controller) GetAllTest(e echo.Context) error {
	token := e.Request().Header.Get(echo.HeaderAuthorization)

	resp, err := ctrl.srv.GetAllTests(e.Request().Context(), token)
	if err != nil {
		ctrl.log.Error(
			"failed to get all tests",
			zap.Error(err),
		)

		return handleErr(err)
	}

	var tests []model.GetTestResponse

	for _, t := range resp {
		test := model.GetTestResponse{
			ID:        t.ID,
			Name:      t.Name,
			Questions: make([]model.TestQuestion, len(resp)),
		}

		for i := 0; i < len(t.Questions); i++ {
			test.Questions[i] = model.TestQuestion{
				Order:    t.Questions[i].Order,
				Question: t.Questions[i].Question,
			}
		}
		tests = append(tests, test)
	}

	ctrl.log.Debug("successfully handle get all tests")

	return e.JSON(http.StatusOK, model.GetAllTestResponse{
		GetTestResponse: tests,
	})
}

func (ctrl *Controller) CreateConsumer(e echo.Context) error {
	var (
		req model.CreateConsumerRequest
	)

	err := e.Bind(&req)
	if err != nil {
		ctrl.log.Error("failed to bind request")
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	resp, err := ctrl.srv.CreateConsumer(e.Request().Context(), dto.CreateConsumer{
		Login:    req.Login,
		Password: req.Password,
	})
	if err != nil {
		ctrl.log.Error("failed to create consumer")
		return handleErr(err)
	}

	ctrl.log.Debug("successfully handle create consumer")

	return e.JSON(http.StatusCreated, model.CreateConsumerResponse{
		AccessToken:  resp.AccessToken,
		RefreshToken: resp.RefreshToken,
	})

}
func (ctrl *Controller) GetConsumer(e echo.Context) error {
	token := e.Request().Header.Get(echo.HeaderAuthorization)

	resp, err := ctrl.srv.GetConsumerByID(e.Request().Context(), token)
	if err != nil {
		ctrl.log.Error("failed to get consumer by id")
		return handleErr(err)
	}

	ctrl.log.Debug("successfully handle get consumer")

	return e.JSON(http.StatusOK, model.GetConsumerResponse{
		ID:        resp.ID,
		Login:     resp.Login,
		CreatedAt: resp.CreatedAt,
	})
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

	ctrl.log.Debug("successfully handle get consumer")

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
