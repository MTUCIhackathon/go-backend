package jwt

import (
	"time"

	"go.uber.org/zap"

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

	prv.log.Debug("created jwt with claims", zap.Any("claims", claims))

	tokenString, err := t.SignedString(prv.privateKey)

	if err != nil {
		prv.log.Debug("failed to sign access token", zap.Error(err))
		return "", token.ErrorSignedToken
	}

	prv.log.Debug("created access token", zap.String("access token", tokenString))

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

	prv.log.Debug("created jwt with claims", zap.Any("claims", claims))

	tokenString, err := t.SignedString(prv.privateKey)

	if err != nil {
		prv.log.Debug("failed to sign refresh token", zap.Error(err))
		return "", token.ErrorSignedToken
	}

	prv.log.Debug("created refresh token", zap.String("refresh token", tokenString))

	return tokenString, nil
}

func (prv *Provider) GetDataFromToken(raw string) (*dto.UserDataInToken, error) {
	prv.log.Debug("start getting data from jwt token")

	jwtToken, err := prv.getJWTFromBearerToken(raw)
	if err != nil {
		prv.log.Debug("failed to parse jwt token", zap.Error(err))
		return nil, err
	}
	parsedToken, err := jwt.ParseWithClaims(jwtToken, &JWT{}, prv.readKeyFunc)

	if err != nil {
		prv.log.Debug("failed to parse jwt token", zap.Error(err))
		return nil, token.ErrorParsedToken
	}

	prv.log.Debug("parsed jwt token", zap.Any("claims", parsedToken))

	claims, ok := parsedToken.Claims.(*JWT)
	if !ok {
		prv.log.Debug("failed to parse jwt token claims")
		return nil, token.ErrorParsedClaims
	}

	if claims.ExpiresAt.Before(time.Now()) {
		prv.log.Debug("expired jwt token", zap.Any("claims", claims))
		return nil, token.ErrorTimeExpired
	}

	prv.log.Debug("parsed jwt token", zap.Any("claims", claims))

	var ParsedID uuid.UUID

	ParsedID, err = uuid.Parse(claims.RegisteredClaims.Issuer)
	if err != nil {
		prv.log.Debug("failed to parse jwt token claim", zap.Error(err))
		return nil, token.ErrorParsedID
	}

	prv.log.Debug("successfully parsed userID", zap.Any("id", ParsedID))

	data := &dto.UserDataInToken{
		ID:       ParsedID,
		IsAccess: claims.IsAccess,
	}

	prv.log.Debug("successfully parsed data", zap.Any("data", data))
	return data, nil
}
