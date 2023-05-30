package api

import (
	"fmt"

	"bitbucket.org/ifan-moladin/base-project/modules/health"
	"bitbucket.org/ifan-moladin/base-project/utils/database"
	"bitbucket.org/ifan-moladin/base-project/utils/environment"
	"bitbucket.org/ifan-moladin/base-project/utils/httpserver"
	_ "github.com/joho/godotenv/autoload"
)

func Default() *Api {
	env := environment.New()
	dbConn, err := database.DefaultMysqlConnectionFromDsn(env.Get("MYSQL_DSN"))
	if err != nil {
		panic(fmt.Errorf("database.DefaultMysqlConnectionFromDsn: %w", err))
	}
	routers := []Router{
		health.NewRouter(dbConn),
	}
	return New(
		httpserver.DefaultGin(),
		dbConn,
		routers,
	)
}
