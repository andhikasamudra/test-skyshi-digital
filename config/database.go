package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"os"
)

func GetConnection() *bun.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?tls=skip-verify&autocommit=true",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_DBNAME"),
	)
	sqldb, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	db := bun.NewDB(sqldb, mysqldialect.New())

	return db
}
