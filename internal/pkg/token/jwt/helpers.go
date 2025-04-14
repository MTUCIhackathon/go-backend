package jwt

import (
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"

	"github.com/MTUCIhackathon/go-backend/internal/pkg/token"
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
