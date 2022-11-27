package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"customers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func showIndexPage(w http.ResponseWriter, r *http.Request) {
	log.Println(log.Ldate, "Displaying index page ...")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.ServeFile(w, r, "./static/index.html")
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	log.Println(log.Ldate, "Displaying details of customer with ID", id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers.GetCustomer(id))
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	log.Println(log.Ldate, "Listing all customers...")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(customers.GetCustomers())
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	log.Println(log.Ldate, "Deleting customer with ID", id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers.DeleteCustomer(id))
}

func main() {
	error := godotenv.Load()
	var port string
	if error != nil {
		port = "3000"
	} else {
		port = os.Getenv("APP_PORT")
	}

	customers.InitializeCustomersDb()
	router := mux.NewRouter()
	router.HandleFunc("/", showIndexPage).Methods("GET")
	router.HandleFunc("/customers", getCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	router.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")

	fmt.Println("Starting server on port", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
	fmt.Println("Server initialized. You may now invoke API requests")

}
