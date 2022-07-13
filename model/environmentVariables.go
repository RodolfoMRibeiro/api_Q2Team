package model

import (
	v6 "github.com/caarlos0/env/v6"
)

type environmentVariables struct {
	Database string `env:"DATABASE"`
	Port     string `env:"MYSQL_PORT" envDefault:"3306"`
	Host     string `env:"MYSQL_HOST"`
	User     string `env:"MYSQL_USER"`
	Password string `env:"MYSQL_PASSWORD"`
}

func (env *environmentVariables) Load() {
	env.setWithEnvFile()
}

func (env *environmentVariables) setWithEnvFile() {
	v6.Parse(env)
}
