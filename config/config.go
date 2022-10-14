package config

import "apollo/pkg/util"

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
	Host:     util.EnvStr("REDIS_HOST", "UNSET"),
	Port:     util.EnvInt("REDIS_PORT", 6379),
	Password: util.EnvStr("REDIS_PASSWORD", ""),
	DB:       util.EnvInt("REDIS_DB", 0),
}
