package hash

import (
	"github.com/MTUCIhackathon/go-backend/internal/pkg/encryptor"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

var _ encrytpor.Interface = (*Encryptor)(nil)

type Option func(encryptor *Encryptor)

type Encryptor struct {
	log                        *zap.Logger
	cost                       int
	generateFromPasswordFunc   func([]byte, int) ([]byte, error)
	compareHashAndPasswordFunc func([]byte, []byte) error
}

func New(log *zap.Logger, opts ...Option) *Encryptor {
	if log == nil {
		log = zap.NewNop()
	}
	e := &Encryptor{
		log:                        log.Named("encryptor"),
		cost:                       bcrypt.DefaultCost,
		generateFromPasswordFunc:   bcrypt.GenerateFromPassword,
		compareHashAndPasswordFunc: bcrypt.CompareHashAndPassword,
	}
	for _, opt := range opts {
		opt(e)
	}
	log.Info("encryptor initialized successfully")
	return e
}

func (e *Encryptor) EncryptPassword(password string) (string, error) {
	hash, err := e.generateFromPasswordFunc([]byte(password), e.cost)
	if err != nil {
		return "", encrytpor.ErrorEncryptPassword
	}
	return string(hash), nil
}

func (e *Encryptor) CompareHashAndPassword(hash, password string) error {
	err := e.compareHashAndPasswordFunc([]byte(hash), []byte(password))
	if err != nil {
		return encrytpor.ErrorDecryptPassword
	}
	return nil
}
