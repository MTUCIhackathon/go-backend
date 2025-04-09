package jwt

import (
	"github.com/MTUCIhackathon/server/internal/models"
	tok "github.com/MTUCIhackathon/server/internal/pkg/token"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

func (prv *Provider) CreateAccessTokenForUser(userID uuid.UUID) (string, error) {
	prv.log.Debug("start creating access token")

	now := time.Now()

	add := time.Duration(prv.accessLifetime) * time.Hour

	claims := &JWT{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    userID.String(),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(add)),
		},
		IsAccess: true,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	tokenString, err := token.SignedString(prv.privateKey)

	if err != nil {
		return "", tok.ErrorSignedToken
	}

	prv.log.Debug("created access token")

	return tokenString, nil
}

func (prv *Provider) CreateRefreshTokenForUser(userID uuid.UUID) (string, error) {
	prv.log.Debug("start creating refresh token")

	now := time.Now()
	add := time.Duration(prv.refreshLifetime) * time.Hour

	claims := &JWT{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    userID.String(),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(add)),
		},
		IsAccess: false,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	tokenString, err := token.SignedString(prv.privateKey)

	if err != nil {
		return "", tok.ErrorSignedToken
	}

	prv.log.Debug("created refresh token")

	return tokenString, nil
}

func (prv *Provider) GetDataFromToken(token string) (*models.UserDataInToken, error) {
	prv.log.Debug("start getting data from jwt token")

	parsedToken, err := jwt.ParseWithClaims(token, &JWT{}, prv.readKeyFunc)

	if err != nil {
		return nil, tok.ErrorParsedToken
	}

	prv.log.Debug("parsed jwt token")

	claims, ok := parsedToken.Claims.(*JWT)
	if !ok {
		return nil, tok.ErrorParsedClaims
	}

	var ParsedID uuid.UUID

	ParsedID, err = uuid.Parse(claims.RegisteredClaims.Issuer)
	if err != nil {
		return nil, tok.ErrorParsedID
	}

	prv.log.Debug("successfully parsed userID")

	data := &models.UserDataInToken{
		UserId:   ParsedID,
		IsAccess: claims.IsAccess,
	}

	return data, nil
}
