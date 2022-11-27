module github.com/rdemirov/crmBackendApi

go 1.19

require (
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
	github.com/rdemirov/crmBackendApi/customers v1.0.0
)

replace github.com/rdemirov/crmBackendApi/customers v1.0.0 => ../customers