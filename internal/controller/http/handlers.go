package http

import (
	"github.com/labstack/echo/v4"

	"github.com/MTUCIhackathon/go-backend/internal/model/dto"
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
	req := new(dto.CreateConsumer)
	
	err := e.Bind(req)
	if err != nil {
		// TODO
		return nil
	}

	return nil
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
