package customer

import (
	"fmt"
	"home/pkg/customer"
	"home/pkg/sqldatabase"
	"log"
)

func getCustomersFromDbById(id int) *customer.Customer {
	rows, err := sqldatabase.DB.Database.Query("SELECT * FROM customer WHERE customerid = $1", id)
	if err != nil {
		log.Println("sql query error: ", err)
		return nil
	}
	defer rows.Close()

	rows.Next()
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
		log.Println(err)
	}
	fmt.Println(sqldatabase.ConstructPostRequest(f))
	return &f
}

func getCustomersFromDb() *[]customer.Customer {
	var Customers []customer.Customer
	rows, err := sqldatabase.DB.Database.Query("SELECT * FROM customer")
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

	return &Customers
}

func addCustomer(cstmer *customer.Customer) error {
	rslt, err := sqldatabase.DB.Database.Exec(
		sqldatabase.ConstructPostRequest(*cstmer),
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
