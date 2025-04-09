package jwt

import (
	"crypto/rsa"
	"fmt"
	"github.com/MTUCIhackathon/server/internal/config"
	"github.com/MTUCIhackathon/server/internal/pkg/token"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"os"
)

var _ token.Provider = (*Provider)(nil)

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
		return nil, fmt.Errorf("failed to read public key: %w", err)
	}

	log.Debug("successful read public key path")

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyRaw)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %w", err)
	}

	log.Debug("successful parse public key")

	privateKeyRaw, err := os.ReadFile(cfg.JWT.PrivateKeyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read private key: %w", err)
	}

	log.Debug("successful read private key path")

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyRaw)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
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
