// Package config
// @Author twilikiss 2024/12/13 18:26:26
package config

type pollConfig struct {
	ServiceConfig ServiceConfig
}

type ServiceConfig struct {
	Url  string
	Time string
}
