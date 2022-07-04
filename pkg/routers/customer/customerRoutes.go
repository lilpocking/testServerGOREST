package customer

import (
	"encoding/json"
	"home/pkg/customer"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func GetCustomerById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

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

func PostCustomer(w http.ResponseWriter, r *http.Request) {
	cstmer := customer.Customer{}
	json.NewDecoder(r.Body).Decode(&cstmer)
	err := addCustomer(&cstmer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}
