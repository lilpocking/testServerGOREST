package dbfunc

import (
	"database/sql"
	"home/config"
	"log"
	"os"
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
		err = db.Close()
		if err != nil {
			log.Println(err)
		}
	}
}
