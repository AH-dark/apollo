package config

type system struct {
	Listen        string `binding:"required"`
	Debug         bool
	SessionSecret string
	HashIDSalt    string
}

type database struct {
	Type        string `binding:"eq=mysql|eq=postgres|eq=sqlite3|eq=mssql"`
	Host        string
	Port        int
	Name        string
	Username    string
	Password    string
	Charset     string
	TablePrefix string
	SSLMode     string `binding:"eq=disable|eq=allow|eq=require|eq=verify-ca|eq=verify-full"` // SSLMode is used for PostgresSQL SSL connections
	FileName    string // FileName is only used for sqlite3
}

type redis struct {
	Network  string
	Host     string
	Port     int
	Password string
	DB       int
}

type cors struct {
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	ExposeHeaders    []string
	AllowCredentials bool
	MaxAge           int
	Https            bool
	SameSite         string
}
