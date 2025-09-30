package model

import "time"

type AppApiConfig struct {
	Loglevel string
	Jwt      Jwt
	Server   Api
	Database Postgres
}
type Postgres struct {
	Host string
	Port string
	User string
	Pwd  string
	Db   string
	Ssl  string
}

type Api struct {
	Host string
	Port string
}

type Jwt struct {
	Secret  string
	Refresh string
	Time    time.Duration
	Long    time.Duration
}
