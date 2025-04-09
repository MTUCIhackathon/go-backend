package jwt

import (
	"crypto/rsa"
	"github.com/MTUCIhackathon/server/internal/config"
	tok "github.com/MTUCIhackathon/server/internal/pkg/token"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"os"
)

var _ tok.Provider = (*Provider)(nil)

type JWT struct {
	jwt.RegisteredClaims
	IsAccess bool `json:"is_access"`
}

type Provider struct {
	log             *zap.Logger
	publicKey       *rsa.PublicKey
	privateKey      *rsa.PrivateKey
	accessLifetime  int
	refreshLifetime int
}

func NewProvider(cfg *config.Config, log *zap.Logger) (*Provider, error) {
	publicKeyRaw, err := os.ReadFile(cfg.JWT.PublicKeyPath)
	if err != nil {
		return nil, tok.ErrorReadPublicKey
	}

	log.Debug("successful read public key path")

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyRaw)
	if err != nil {
		return nil, tok.ErrorParsedPublicKey
	}

	log.Debug("successful parse public key")

	privateKeyRaw, err := os.ReadFile(cfg.JWT.PrivateKeyPath)
	if err != nil {
		return nil, tok.ErrorReadPrivateKey
	}

	log.Debug("successful read private key path")

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyRaw)
	if err != nil {
		return nil, tok.ErrorParsedPrivateKey
	}

	log.Debug("successful parse private key")
	provider := &Provider{
		log:             log,
		publicKey:       publicKey,
		privateKey:      privateKey,
		accessLifetime:  cfg.JWT.AccessTokenLifeTime,
		refreshLifetime: cfg.JWT.RefreshTokenLifeTime,
	}
	log.Debug("successful create new jwt provider")
	return provider, nil
}
