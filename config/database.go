package config

import ("./database")

var Database = map[string]interface{} {
	"driver": "mysql",
	"mysql": database.Mysql,
}
