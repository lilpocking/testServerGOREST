package sqldatabase

import (
	"database/sql"
	"home/config"
	"home/pkg/dboprtns"
)

type dataBase struct {
	Database   *sql.DB
	Is_Started bool
}

var DB dataBase

func Open() (err error) {
	dboprtns.CheckDbFileExist() // Проверка существует ли SQLite база данных, если не существует, то создает файл и нужные таблицы
	DB.Database, err = sql.Open(config.DriverName, config.DbPath)
	if err != nil {
		DB.Is_Started = false
	} else {
		DB.Is_Started = true
	}
	return
}

func Close() (err error) {
	if DB.Database != nil && DB.Is_Started {
		err = DB.Database.Close()
	}
	return
}
