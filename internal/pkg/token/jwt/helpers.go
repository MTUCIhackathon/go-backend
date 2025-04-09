package jwt

import (
	tok "github.com/MTUCIhackathon/server/internal/pkg/token"
	"github.com/golang-jwt/jwt/v5"
)

func (prv *Provider) readKeyFunc(token *jwt.Token) (interface{}, error) {
	switch token.Method.(type) {
	case *jwt.SigningMethodRSA:
		prv.log.Debug("successful read key")
		return prv.publicKey, nil
	default:
		return nil, tok.ErrorMethod
	}
}
