package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"strconv"
)

func Default() (*gorm.DB, error) {
	// logLevel (env: DB_LOG_MODE)
	// 1 = Silent (not printing any log)
	// 2 = Error (only printing in case of error)
	// 3 = Warn (print error + warning)
	// 4 = Info (print all type of log)
	logLevel, _ := strconv.Atoi(os.Getenv("DB_LOG_MODE"))
	if logLevel == 0 {
		logLevel = 2
	}

	connString := os.Getenv("MYSQL_CONNECTION")
	dbConn, err := gorm.Open(
		mysql.Open(connString),
		&gorm.Config{
			CreateBatchSize: 500,
			Logger:          logger.Default.LogMode(logger.LogLevel(logLevel)),
		},
	)

	return dbConn, err

}
