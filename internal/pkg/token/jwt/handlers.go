package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"github.com/MTUCIhackathon/go-backend/internal/model/dto"
	"github.com/MTUCIhackathon/go-backend/internal/pkg/token"
)

func (prv *Provider) CreateAccessTokenForUser(userID uuid.UUID) (string, error) {
	prv.log.Debug("start creating access token")

	now := time.Now()

	add := time.Duration(prv.accessLifeTime) * time.Hour

	claims := &JWT{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    userID.String(),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(add)),
		},
		IsAccess: true,
	}

	t := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	tokenString, err := t.SignedString(prv.privateKey)

	if err != nil {
		return "", token.ErrorSignedToken
	}

	prv.log.Debug("created access token")

	return tokenString, nil
}

func (prv *Provider) CreateRefreshTokenForUser(userID uuid.UUID) (string, error) {
	prv.log.Debug("start creating refresh token")

	now := time.Now()
	add := time.Duration(prv.refreshLifeTime) * time.Hour

	claims := &JWT{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    userID.String(),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(add)),
		},
		IsAccess: false,
	}

	t := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	tokenString, err := t.SignedString(prv.privateKey)

	if err != nil {
		return "", token.ErrorSignedToken
	}

	prv.log.Debug("created refresh token")

	return tokenString, nil
}

func (prv *Provider) GetDataFromToken(jwtToken string) (*dto.UserDataInToken, error) {
	prv.log.Debug("start getting data from jwt token")

	parsedToken, err := jwt.ParseWithClaims(jwtToken, &JWT{}, prv.readKeyFunc)

	if err != nil {
		return nil, token.ErrorParsedToken
	}

	prv.log.Debug("parsed jwt token")

	claims, ok := parsedToken.Claims.(*JWT)
	if !ok {
		return nil, token.ErrorParsedClaims
	}

	var ParsedID uuid.UUID

	ParsedID, err = uuid.Parse(claims.RegisteredClaims.Issuer)
	if err != nil {
		return nil, token.ErrorParsedID
	}

	prv.log.Debug("successfully parsed userID")

	data := &dto.UserDataInToken{
		UserID:   ParsedID,
		IsAccess: claims.IsAccess,
	}

	return data, nil
}
