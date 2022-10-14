package model

import (
	"fmt"
	"github.com/AH-dark/apollo/config"
	"github.com/AH-dark/apollo/pkg/log"
	"github.com/AH-dark/apollo/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	gorm_logger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// initDB initializes the database and migrates the schema.
func initDB(forceMigrate bool) *gorm.DB {
	var dialector gorm.Dialector

	if gin.Mode() == gin.TestMode {
		dialector = sqlite.Open("file::memory:?cache=shared")
	} else {
		switch config.Database.Type {
		case "sqlite", "sqlite3":
			dialector = sqlite.Open(util.AbsolutePath(config.Database.FileName))
		case "mysql", "mariadb":
			dialector = mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
				config.Database.Username,
				config.Database.Password,
				config.Database.Host,
				config.Database.Port,
				config.Database.Name,
				config.Database.Charset,
			))
		case "postgres", "postgresql":
			dialector = postgres.Open(fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s TimeZone=Asia/Shanghai",
				config.Database.Host,
				config.Database.Port,
				config.Database.Username,
				config.Database.Name,
				config.Database.Password,
				config.Database.SSLMode,
			))
		case "mssql", "sqlserver":
			dialector = sqlserver.Open(fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
				config.Database.Username,
				config.Database.Password,
				config.Database.Host,
				config.Database.Port,
				config.Database.Name,
			))
		default:
			log.Log().Panicf("database type %s not support", config.Database.Type)
		}
	}

	db, err := gorm.Open(dialector, &gorm.Config{
		PrepareStmt: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			TablePrefix:   config.Database.TablePrefix,
		},
		Logger: gorm_logger.New(log.Log(), gorm_logger.Config{
			LogLevel: lo.If(config.System.Debug, gorm_logger.Info).Else(gorm_logger.Silent),
		}),
	})
	if err != nil {
		log.Log().WithError(err).Panic("database connect error")
		return nil
	}

	migrate(db, forceMigrate)

	return db
}
