package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rdemirov/crmBackendApi/customers"
)

func showIndexPage(w http.ResponseWriter, r *http.Request) {
	log.Println(log.Ldate, "Displaying index page ...")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.ServeFile(w, r, "./static/index.html")
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

	fmt.Println("Starting server on port", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
	fmt.Println("Server initialized. You may now invoke API requests")

}
