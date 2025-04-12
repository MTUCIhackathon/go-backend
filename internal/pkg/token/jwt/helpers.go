package jwt

import (
	"github.com/MTUCIhackathon/go-backend/internal/pkg/token"
	"github.com/golang-jwt/jwt/v5"
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
