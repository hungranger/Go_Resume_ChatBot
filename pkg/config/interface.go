package config

type IConfig interface {
	LoadConfig() IConfig
	GetPort() string
	GetFbApi() string
	GetImage() string
}

func ProvideIniConfig() IConfig {
	return new(IniConfig).LoadConfig()
}
