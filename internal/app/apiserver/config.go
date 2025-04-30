package apiserver

import "webgo/internal/app/store"

// Config ...
type Config struct {
	BindAddr string `toml:"bind_addr"` //адрес на котором запускаем веб сервер
	LogLevel string `toml:"log_level"`
	Store    *store.Config
}

//NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr: "8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
