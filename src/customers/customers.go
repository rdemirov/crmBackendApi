package customers

import (
	"fmt"
	"log"
	"strconv"

	"github.com/google/uuid"
)

type Customer struct {
	Id        string
	Name      string
	Role      string
	Email     string
	Phone     string
	Contacted bool
}

func (c *Customer) update(customerData map[string]string) {
	fmt.Println(customerData)
	if customerData["Name"] != "" {
		c.Name = customerData["Name"]
	}
	if customerData["Role"] != "" {
		c.Role = customerData["Role"]
	}
	if customerData["Email"] != "" {
		c.Email = customerData["Email"]
	}
	if customerData["Phone"] != "" {
		c.Phone = customerData["Phone"]
	}
	if customerData["Contacted"] != "" {
		c.Contacted, _ = strconv.ParseBool(customerData["Contacted"])
	}

	fmt.Println(c)

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

func UpdateCustomer(id string, customerData map[string]string) bool {
	customerToUpdate := customers[id]
	if customerToUpdate.Id == "" {
		return false
	}
	customerToUpdate.update(customerData)
	customers[id] = customerToUpdate
	return true
}

func DeleteCustomer(id string) bool {
	customerToDelete := customers[id]
	if customerToDelete.Id == "" {
		return false
	}
	delete(customers, id)
	return true
}

func AddCustomer(c Customer) string {
	customerUuid := uuid.New()
	customerId := customerUuid.String()
	newCustomer := Customer{
		Id:        customerId,
		Name:      c.Name,
		Phone:     c.Phone,
		Email:     c.Email,
		Role:      c.Role,
		Contacted: c.Contacted,
	}
	customers[customerId] = newCustomer
	return customerId
}
