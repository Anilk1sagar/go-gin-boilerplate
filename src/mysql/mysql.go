package mysql

import (
	"gitlab.com/anilk1sagar/go_gin_test/src/utils"
)

// MysqlConfig for mysqlCnfig struct
type MysqlConfig struct {
	MysqlDb *DBConfig
}

// DBConfig for mysqldb struct
type DBConfig struct {
	Dialect  string
	Username string
	Password string
	Name     string
	Charset  string
}

// GetMysqlConfig for mysqldb config
func GetMysqlConfig() *MysqlConfig {

	return &MysqlConfig{

		MysqlDb: &DBConfig{
			Dialect:  "mysql",
			Username: utils.MysqlUsername(),
			Password: utils.MysqlPassword(),
			Name:     utils.MysqlDBName(),
			Charset:  "utf8",
		},
	}
}
