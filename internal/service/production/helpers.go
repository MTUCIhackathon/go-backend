package production

import (
	"errors"
	"fmt"
	"github.com/MTUCIhackathon/go-backend/internal/model/dto"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

type handleError struct{}

func (s *Service) getConsumerIDFromRequest(req *http.Request) (uuid.UUID, error) {
	var (
		data *dto.UserDataInToken
		err  error
	)
	data, err = s.getConsumerDataFromRequest(req)
	if err != nil {
		s.log.Debug("failed to get consumer data from request", zap.Error(err))
		return uuid.Nil, err
	}

	if !data.IsAccess {
		s.log.Debug("failed to get consumer data from request", zap.Any("data", data))
		return uuid.Nil, errors.New("not authorized")
	}

	s.log.Debug("got consumer data from request", zap.Any("data", data))

	return data.ID, nil
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
		s.log.Debug("failed to get JWT", zap.Error(err))
		return nil, err
	}

	s.log.Debug("got JWT", zap.Any("token", token))

	userData, err = s.provider.GetDataFromToken(token)
	if err != nil {
		s.log.Debug("failed to get data from token", zap.Error(err))
		return nil, err
	}

	s.log.Debug("got data from token", zap.Any("userData", userData))

	return userData, nil
}

func (s *Service) getJWTFromBearerToken(raw string) (string, error) {
	splitToken := strings.Split(raw, "Bearer")
	if len(splitToken) != 2 {
		s.log.Debug("failed to parse bearer token", zap.Any("token", splitToken))
		return "", fmt.Errorf("invalid token")
	}

	s.log.Debug("got bearer token", zap.Any("token", splitToken))

	reqToken := strings.TrimSpace(splitToken[1])

	s.log.Debug("got jwt token", zap.Any("token", reqToken))

	return reqToken, nil
}

func (s *Service) validRefreshToken(req *http.Request) (uuid.UUID, error) {
	var (
		data *dto.UserDataInToken
		err  error
	)

	data, err = s.getConsumerDataFromRequest(req)
	if err != nil {
		s.log.Debug("failed to get consumer data from request", zap.Error(err))
		return uuid.Nil, fmt.Errorf("invalid token")
	}

	s.log.Debug("got consumer data from request", zap.Any("data", data))

	if data.IsAccess {
		s.log.Debug("token is not a refresh token", zap.Any("token", data))
		return uuid.Nil, fmt.Errorf("token is not a refresh token")
	}

	s.log.Debug("got token from request", zap.Any("token", data))

	return data.ID, nil
}
