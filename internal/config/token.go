package config

type Token struct {
	AccessTokenLifeTime  int    `config:"access_token_life_time" toml:"access_token_life_time" yaml:"access_token_life_time" json:"access_token_life_time"`
	RefreshTokenLifeTime int    `config:"refresh_token_life_time" toml:"refresh_token_life_time" yaml:"refresh_token_life_time" json:"refresh_token_life_time"`
	PublicKeyPath        string `config:"public_key_path" toml:"public_key_path" yaml:"public_key_path" json:"public_key_path"`
	PrivateKeyPath       string `config:"private_key_path" toml:"private_key_path" yaml:"private_key_path" json:"private_key_path"`
}
