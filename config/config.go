package config

import "github.com/AH-dark/apollo/pkg/util"

const (
	AppVersion = "beta.1.0.0"
)

var System = &system{
	Listen:        util.EnvStr("LISTEN", ":8080"),
	Debug:         util.EnvBool("DEBUG", false),
	SessionSecret: util.EnvStr("SESSION_SECRET", util.RandString(32)),
	HashIDSalt:    util.EnvStr("HASHID_SALT", util.RandString(32)),
}

var Database = &database{
	Type:        util.EnvStr("DB_TYPE", "sqlite3"),
	Host:        util.EnvStr("DB_HOST", "localhost"),
	Port:        util.EnvInt("DB_PORT", 3306),
	Name:        util.EnvStr("DB_NAME", "apollo"),
	Username:    util.EnvStr("DB_USERNAME", "root"),
	Password:    util.EnvStr("DB_PASSWORD", ""),
	Charset:     util.EnvStr("DB_CHARSET", "utf8"),
	TablePrefix: util.EnvStr("DB_PREFIX", "apollo_"),
	SSLMode:     util.EnvStr("DB_SSL", "disable"),
	FileName:    util.EnvStr("DB_FILE", "apollo.db"),
}

var Redis = &redis{
	Network:  util.EnvStr("REDIS_NETWORK", "tcp"),
	Host:     util.EnvStr("REDIS_HOST", ""),
	Port:     util.EnvInt("REDIS_PORT", 6379),
	Password: util.EnvStr("REDIS_PASSWORD", ""),
	DB:       util.EnvInt("REDIS_DB", 0),
}

var CORS = &cors{
	AllowOrigins:     util.EnvStringSlice("CORS_ALLOW_ORIGINS", ",", []string{"*"}),
	AllowMethods:     util.EnvStringSlice("CORS_ALLOW_METHODS", ",", []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"}),
	AllowHeaders:     util.EnvStringSlice("CORS_ALLOW_HEADERS", ",", []string{"Origin", "Content-Length", "Content-Type", "Authorization"}),
	ExposeHeaders:    util.EnvStringSlice("CORS_EXPOSE_HEADERS", ",", []string{}),
	MaxAge:           util.EnvInt("CORS_MAX_AGE", 86400),
	AllowCredentials: util.EnvBool("CORS_ALLOW_CREDENTIALS", true),
}
