package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	customer := customers.GetCustomer(id)
	if customer.Id == "" {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(nil)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(customer)
	}
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
	deletionResult := customers.DeleteCustomer(id)

	if !deletionResult {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	json.NewEncoder(w).Encode(deletionResult)
}

func addCustomer(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var customer customers.Customer
	err := decoder.Decode(&customer)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	log.Println(log.Ldate, "Creating New Customer", customer.Id)
	creationResult := customers.AddCustomer(customer)

	if creationResult == "" {
		w.WriteHeader(http.StatusInternalServerError)

	} else {
		w.WriteHeader(http.StatusCreated)
	}

	json.NewEncoder(w).Encode(creationResult)
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var customerMap = make(map[string]string)
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &customerMap)
	w.Header().Set("Content-Type", "application/json")
	log.Println(log.Ldate, "Updating Customer", id)
	updateResult := customers.UpdateCustomer(id, customerMap)
	if !updateResult {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	json.NewEncoder(w).Encode(updateResult)
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
	router.HandleFunc("/customers", addCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PATCH")
	router.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")

	fmt.Println("Starting server on port", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
	fmt.Println("Server initialized. You may now invoke API requests")

}
