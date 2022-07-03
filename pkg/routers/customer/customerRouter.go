package customer

import (
	"database/sql"
	"encoding/json"
	"home/config"
	"home/pkg/customer"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func GetCustomerById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	//CORS устанавливается в mainRouter.go
	//w.Header().Set("Access-Control-Allow-Origin", "*") // длы решения проблемы с CORS политикой
	w.Header().Set("Content-Type", "application/json") // установка типа отправляемого контента

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
	}

	Customer := getCustomersFromDbById(id)

	json.NewEncoder(w).Encode(Customer) // кодирование структуры в json формат
}

func GetCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	customers := getCustomersFromDb()

	json.NewEncoder(w).Encode(customers)
}

func getCustomersFromDbById(id int) (Customer customer.Customer) {
	db, err := sql.Open(config.DriverName, config.DbPath)
	if err != nil {
		log.Println("sql open error: ", err)
	}

	rows, err := db.Query("SELECT * FROM customer WHERE customerid = $1", id)
	if err != nil {
		log.Println("sql error: ", err)
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
			log.Println(err)
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

//func for get customer from db
