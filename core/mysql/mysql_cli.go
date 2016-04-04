package mysql

import (
	 "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

var instanceMysqlCli *sql.DB = nil

func Connect() (cn *sql.DB) {
	if instanceMysqlCli == nil {
		var err error

		instanceMysqlCli, err = sql.Open("mysql", "root:354216@/hoditgo")

		if err != nil {
			panic(err)
		}
	}

	return instanceMysqlCli
}
