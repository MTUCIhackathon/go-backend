package validator

type Interface interface {
	ValidatePassword(password string) error
	ValidateEmail(email string) error
}
