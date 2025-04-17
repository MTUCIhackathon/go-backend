package client

import (
	"fmt"
	"net/smtp"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var (
	ErrZeroLengthProfessions   = errors.New("received zero length professions")
	ErrInvalidProfessionLength = errors.New("unexpected profession length: should be equal to 3")
)

func (s *SMTP) SendResultOnEmail(professions []string, testName string, email string) error {
	if len(professions) == 0 {
		s.log.Error(
			"got nil professions list: should be more than zero",
			zap.Error(ErrZeroLengthProfessions),
		)

		return ErrZeroLengthProfessions
	}

	if len(professions) != 3 {
		s.log.Error(
			"bad length of professions",
			zap.Int("professions length", len(professions)),
		)

		return ErrInvalidProfessionLength
	}

	sprintedBytes := []byte(fmt.Sprintf(html, testName, professions[0], professions[1], professions[2]))

	smtpAddress := s.cfg.SMTP.GetSMTPServerAddress()

	plainAuth := smtp.PlainAuth("", s.cfg.SMTP.Login, s.cfg.SMTP.Password, smtpAddress)

	s.log.Debug("created smtp plain auth")

	err := smtp.SendMail(smtpAddress, plainAuth, s.cfg.SMTP.Login, []string{email}, sprintedBytes)
	if err != nil {
		s.log.Error(
			"failed to send mail",
			zap.Error(err),
			zap.String("email", email),
		)

		return errors.Wrap(err, "failed to send mail")
	}

	s.log.Debug("successfully sent email", zap.String("email", email))

	return nil
}
