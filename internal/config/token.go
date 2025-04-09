package config

type Token struct {
	AccessTokenLifeTime  int    `config:"access_token_life_time"`
	RefreshTokenLifeTime int    `config:"refresh_token_life_time"`
	PublicKeyPath        string `config:"public_key_path"`
	PrivateKeyPath       string `config:"private_key_path"`
}
