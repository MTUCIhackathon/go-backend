package jwt

import (
	"errors"
	"fmt"
	"github.com/MTUCIhackathon/go-backend/internal/model/dto"
	"github.com/MTUCIhackathon/go-backend/internal/pkg/token"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"strings"
)

func (prv *Provider) readKeyFunc(t *jwt.Token) (interface{}, error) {
	switch t.Method.(type) {
	case *jwt.SigningMethodRSA:
		prv.log.Debug("successful read key")
		return prv.publicKey, nil
	default:
		prv.log.Debug("unsupported signing method")
		return nil, token.ErrorMethod
	}
}

func (prv *Provider) GetConsumerIDFromRequest(token string) (uuid.UUID, error) {
	var (
		data *dto.UserDataInToken
		err  error
	)
	data, err = prv.getConsumerDataFromRequest(token)
	if err != nil {
		prv.log.Debug("failed to get consumer data from request", zap.Error(err))
		return uuid.Nil, err
	}

	if !data.IsAccess {
		prv.log.Debug("failed to get consumer data from request", zap.Any("data", data))
		return uuid.Nil, errors.New("not authorized")
	}

	prv.log.Debug("got consumer data from request", zap.Any("data", data))

	return data.ID, nil
}

func (prv *Provider) getConsumerDataFromRequest(token string) (*dto.UserDataInToken, error) {
	var (
		data     string
		err      error
		userData *dto.UserDataInToken
	)
	token, err = prv.getJWTFromBearerToken(data)
	if err != nil {
		prv.log.Debug("failed to get JWT", zap.Error(err))
		return nil, err
	}

	prv.log.Debug("got JWT", zap.Any("token", token))

	userData, err = prv.GetDataFromToken(token)
	if err != nil {
		prv.log.Debug("failed to get data from token", zap.Error(err))
		return nil, err
	}

	prv.log.Debug("got data from token", zap.Any("userData", userData))

	return userData, nil
}

func (prv *Provider) getJWTFromBearerToken(raw string) (string, error) {
	splitToken := strings.Split(raw, "Bearer")
	if len(splitToken) != 2 {
		prv.log.Debug("failed to parse bearer token", zap.Any("token", splitToken))
		return "", fmt.Errorf("invalid token")
	}

	prv.log.Debug("got bearer token", zap.Any("token", splitToken))

	reqToken := strings.TrimSpace(splitToken[1])

	prv.log.Debug("got jwt token", zap.Any("token", reqToken))

	return reqToken, nil
}

func (prv *Provider) ValidRefreshToken(token string) (uuid.UUID, error) {
	var (
		data *dto.UserDataInToken
		err  error
	)

	data, err = prv.getConsumerDataFromRequest(token)
	if err != nil {
		prv.log.Debug("failed to get consumer data from request", zap.Error(err))
		return uuid.Nil, fmt.Errorf("invalid token")
	}

	prv.log.Debug("got consumer data from request", zap.Any("data", data))

	if data.IsAccess {
		prv.log.Debug("token is not a refresh token", zap.Any("token", data))
		return uuid.Nil, fmt.Errorf("token is not a refresh token")
	}

	prv.log.Debug("got token from request", zap.Any("token", data))

	return data.ID, nil
}
