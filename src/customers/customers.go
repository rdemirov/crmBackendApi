package customers

import "log"

type Customer struct {
	Id        string
	Name      string
	Role      string
	Email     string
	Phone     string
	Contacted bool
}

var customers = make(map[string]Customer)

func InitializeCustomersDb() {
	customers["1"] = Customer{
		Id:        "1",
		Name:      "Johaness Cabal",
		Role:      "Administrator",
		Email:     "jcabal@gmail.com",
		Phone:     "052262928",
		Contacted: true,
	}
	customers["2"] = Customer{
		Id:        "2",
		Name:      "Horst Cabal",
		Role:      "CEO",
		Email:     "hcabal@gmail.com",
		Phone:     "0887848118",
		Contacted: false,
	}
	customers["3"] = Customer{
		Id:        "3",
		Name:      "Max Payne",
		Role:      "Head of Security",
		Email:     "mpayne@gmail.com",
		Phone:     "0887848228",
		Contacted: true,
	}
	log.Println(log.Ldate, "Customer DB Initialized")
}

func GetCustomers() []Customer {
	customersList := []Customer{}
	for _, value := range customers {
		customersList = append(customersList, value)
	}
	return customersList
}

func GetCustomer(id string) Customer {
	return customers[id]
}

func DeleteCustomer(id string) bool {
	delete(customers, id)
	return true
}

func AddCustomer(id string, name string, phone string, email string, role string, contacted bool) string {
	newCustomer := Customer{
		Id:        id,
		Name:      name,
		Phone:     phone,
		Email:     email,
		Role:      role,
		Contacted: contacted,
	}
	customers[id] = newCustomer
	return id
}
