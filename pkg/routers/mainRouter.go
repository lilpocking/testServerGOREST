package routers

import (
	"home/pkg/routers/customer"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateRouter() (r *mux.Router) {
	r = mux.NewRouter()

	r.Use(commonMiddleware)

	//FUNDS
	//Get
	r.HandleFunc("/customers/id:{id:[0-9]+}", customer.GetCustomerById).Methods("GET")
	r.HandleFunc("/customers", customer.GetCustomers).Methods("GET")

	return
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.Method+": ", r.URL.Path)

			w.Header().Add("Access-Control-Allow-Origin", "*")
			next.ServeHTTP(w, r)
		},
	)
}
