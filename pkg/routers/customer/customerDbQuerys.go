package customer

import (
	"database/sql"
	"fmt"
	"home/config"
	"home/pkg/customer"
	"log"
)

func getCustomersFromDbById(id int) (Customer customer.Customer) {

	db, err := sql.Open(config.DriverName, config.DbPath)
	if err != nil {
		log.Println("sql open error: ", err)
	}

	rows, err := db.Query("SELECT * FROM customer WHERE customerid = $1", id)
	if err != nil {
		log.Println("sql query error: ", err)
	}

	for rows.Next() {
		f := customer.Customer{}
		err = rows.Scan(
			&f.Id,
			&f.FirstName,
			&f.LastName,
			&f.Company,
			&f.Address,
			&f.City,
			&f.State,
			&f.Country,
			&f.PostalCode,
			&f.Phone,
			&f.Fax,
			&f.Email,
			&f.SupportRepId)

		if err != nil {
			log.Println("sql read error: ", err)
		}

		Customer = f
	}
	return
}

func getCustomersFromDb() (Customers []customer.Customer) {
	db, err := sql.Open(config.DriverName, config.DbPath)
	if err != nil {
		log.Println(err)
	}

	rows, err := db.Query("SELECT * FROM customer")
	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		f := customer.Customer{}

		err := rows.Scan(
			&f.Id,
			&f.FirstName,
			&f.LastName,
			&f.Company,
			&f.Address,
			&f.City,
			&f.State,
			&f.Country,
			&f.PostalCode,
			&f.Phone,
			&f.Fax,
			&f.Email,
			&f.SupportRepId)

		if err != nil {
			log.Println(err)
		}

		Customers = append(Customers, f)
	}

	return
}

func addCustomer(cstmer *customer.Customer) error {
	var db *sql.DB
	db, err := sql.Open(config.DriverName, config.DbPath)
	if err != nil {
		log.Println("sql post error: ", err)
	}
	rslt, err := db.Exec(
		"INSERT INTO customer ("+
			"FirstName, "+
			"LastName, "+
			"Company, "+
			"Address, "+
			"City, "+
			"State, "+
			"Country, "+
			"PostalCode, "+
			"Phone, "+
			"Fax, "+
			"Email, "+
			"SupportRepId"+
			") VALUES ("+
			"$1, "+
			"$2, "+
			"$3, "+
			"$4, "+
			"$5, "+
			"$6, "+
			"$7, "+
			"$8, "+
			"$9, "+
			"$10, "+
			"$11, "+
			"$12"+
			");",
		cstmer.FirstName,
		cstmer.LastName,
		cstmer.Company,
		cstmer.Address,
		cstmer.City,
		cstmer.State,
		cstmer.Country,
		cstmer.PostalCode,
		cstmer.Phone,
		cstmer.Fax,
		cstmer.Email,
		cstmer.SupportRepId)

	if err != nil {
		log.Println("sql post error: ", err, fmt.Sprintf(" rslt: %v\n", rslt))
	}
	return err
}
