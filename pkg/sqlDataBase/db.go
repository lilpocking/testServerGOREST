package sqldatabase

import (
	"database/sql"
	"home/config"
	"home/pkg/dbfunc"
	"reflect"
	"strconv"
)

type dataBase struct {
	Database   *sql.DB
	Is_Started bool
}

var DB dataBase

func Open() (err error) {
	dbfunc.CheckDbFileExist() // Проверка существует ли SQLite база данных, если не существует, то создает файл и нужные таблицы
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

func ConstructPostRequest(object interface{}) string {
	exQueryMessage := "INSERT INTO customer ("
	t := reflect.TypeOf(object)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if field.Name == "Id" {
			continue
		}

		if i >= t.NumField()-1 {
			exQueryMessage += field.Name + ") VALUES ("
			break
		}
		exQueryMessage += field.Name + ", "
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Name == "Id" {
			continue
		}
		if i >= t.NumField()-1 {
			exQueryMessage += "$" + strconv.Itoa(i+1) + ");"
			break
		}
		exQueryMessage += "$" + strconv.Itoa(i+1) + ", "
	}

	return exQueryMessage
}
