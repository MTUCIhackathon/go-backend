package config

const (
	SigningAlgorithmRS256 string = "RS256"
)

type Token struct {
	AccessTokenLifeTime  int    `config:"access_token_life_time" toml:"access_token_life_time" yaml:"access_token_life_time" json:"access_token_life_time"`
	RefreshTokenLifeTime int    `config:"refresh_token_life_time" toml:"refresh_token_life_time" yaml:"refresh_token_life_time" json:"refresh_token_life_time"`
	PublicKeyPath        string `config:"public_key_path" toml:"public_key_path" yaml:"public_key_path" json:"public_key_path"`
	PrivateKeyPath       string `config:"private_key_path" toml:"private_key_path" yaml:"private_key_path" json:"private_key_path"`
	SigningAlgorithm     string `config:"signing_algorithm" toml:"signing_algorithm" yaml:"signing_algorithm" json:"signing_algorithm"`
}

func (t *Token) copy() *Token {
	if t == nil {
		return nil
	}

	return &Token{
		AccessTokenLifeTime:  t.AccessTokenLifeTime,
		RefreshTokenLifeTime: t.RefreshTokenLifeTime,
		PublicKeyPath:        t.PublicKeyPath,
		PrivateKeyPath:       t.PrivateKeyPath,
		SigningAlgorithm:     t.SigningAlgorithm,
	}
}
