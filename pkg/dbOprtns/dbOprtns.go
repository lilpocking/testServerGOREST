package dboprtns

import (
	"database/sql"
	"home/config"
	"log"
	"os"
	"reflect"
	"strconv"
)

func CheckDbFileExist() {
	var db *sql.DB

	if _, err := os.Stat("Chinook.db"); os.IsNotExist(err) {
		log.Println("Creating db and table customer")
		db, _ = sql.Open(config.DriverName, config.DbPath)

		r, err := db.Exec(
			"CREATE TABLE customer (" +
				"CustomerId INTEGER PRIMARY KEY AUTOINCREMENT, " +
				"FirstName TINYTEXT DEFAULT \"\", " +
				"LastName TINYTEXT DEFAULT \"\", " +
				"Company TINYTEXT DEFAULT \"\", " +
				"Address TINYTEXT DEFAULT \"\", " +
				"City TINYTEXT DEFAULT \"\", " +
				"State TINYTEXT DEFAULT \"\", " +
				"Country TINYTEXT DEFAULT \"\", " +
				"PostalCode TINYTEXT DEFAULT \"\", " +
				"Phone TINYTEXT DEFAULT \"\", " +
				"Fax TINYTEXT DEFAULT \"\", " +
				"Email TINYTEXT DEFAULT \"\", " +
				"SupportRepId INT DEFAULT 0" +
				");")

		log.Println("SQL creating result: ", &r)
		if err != nil {
			log.Println("SQL create error: ", err)
		}

		log.Println("SQL end creating")

		if err = db.Close(); err != nil {
			log.Println(err)
		}
	}
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
