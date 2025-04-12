package production

import (
	"errors"
	"fmt"
	"github.com/MTUCIhackathon/go-backend/internal/model/dto"
	"github.com/google/uuid"
	"net/http"
	"strings"
)

type handleError struct{}

func (s *Service) getConsumerIDFromRequest(req *http.Request) (uuid.UUID, error) {
	var (
		data *dto.UserDataInToken
		ID   uuid.UUID
		err  error
	)
	data, err = s.getConsumerDataFromRequest(req)
	if err != nil {
		return uuid.Nil, err
	}

	if !data.IsAccess {
		return uuid.Nil, errors.New("not authorized")
	}

	return ID, nil
}

func (s *Service) getConsumerDataFromRequest(req *http.Request) (*dto.UserDataInToken, error) {
	var (
		data     string
		token    string
		err      error
		userData *dto.UserDataInToken
	)

	data = req.Header.Get("Authorization")

	token, err = s.getJWTFromBearerToken(data)
	if err != nil {
		return nil, err
	}
	userData, err = s.provider.GetDataFromToken(token)
	if err != nil {
		return nil, err
	}
	return userData, nil
}

func (s *Service) getJWTFromBearerToken(raw string) (string, error) {
	splitToken := strings.Split(raw, "Bearer")
	if len(splitToken) != 2 {
		return "", fmt.Errorf("invalid token")
	}

	reqToken := strings.TrimSpace(splitToken[1])

	return reqToken, nil
}

func (s *Service) validRefreshToken(req *http.Request) (uuid.UUID, error) {
	var (
		data *dto.UserDataInToken
		err  error
	)

	data, err = s.getConsumerDataFromRequest(req)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid token")
	}

	if data.IsAccess {
		return uuid.Nil, fmt.Errorf("token is not a refresh token")
	}
	return data.ID, nil
}
