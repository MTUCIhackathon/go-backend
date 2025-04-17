package smtp

type Interface interface {
	SendResultOnEmail(professions []string, testType string, email string) error
}
