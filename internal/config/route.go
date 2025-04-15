package config

// TODO add copy func and struct tags
type Route struct {
	FirstRoute     string `config:"first_test" toml:"first_test"`
	SecondRoute    string `config:"second_test"`
	ThirdRoute     string `config:"third_test"`
	SummarizeRoute string `config:"summarize_test"`
}
