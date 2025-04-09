package jwt

import (
	"fmt"
	"github.com/MTUCIhackathon/server/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

func (prv *Provider) CreateTokenForUser(userID uuid.UUID, isAccess bool) (string, error) {
	prv.log.Debug("start creating jwt token")

	now := time.Now()
	var add time.Duration

	if isAccess {
		add = time.Duration(prv.accessLifetime) * time.Hour
	} else {
		add = time.Duration(prv.refreshLifetime) * time.Hour
	}

	claims := &JWT{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    userID.String(),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(add)),
		},
		IsAccess: isAccess,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	tokenString, err := token.SignedString(prv.privateKey)

	if err != nil {
		return "", err
	}

	prv.log.Debug("created jwt token")

	return tokenString, nil
}

func (prv *Provider) GetDataFromToken(token string) (*models.UserDataInToken, error) {
	prv.log.Debug("start getting data from jwt token")

	parsedToken, err := jwt.ParseWithClaims(token, &JWT{}, prv.readKeyFunc)

	if err != nil {
		return nil, err
	}

	prv.log.Debug("parsed jwt token")

	claims, ok := parsedToken.Claims.(*JWT)
	if !ok {
		return nil, fmt.Errorf("failed to parse jwt token: invalid claims")
	}

	var ParsedID uuid.UUID

	ParsedID, err = uuid.Parse(claims.RegisteredClaims.Issuer)
	if err != nil {
		return nil, err
	}

	prv.log.Debug("successfully parsed userID")

	data := &models.UserDataInToken{
		UserId:   ParsedID,
		IsAccess: claims.IsAccess,
	}

	return data, nil
}
