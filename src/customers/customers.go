package customers

type Customer struct {
	id        string
	name      string
	role      string
	email     string
	phone     string
	contacted bool
}

var customers map[string]Customer

func InitializeCustomersDb() {
	customers["1"] = Customer{
		id:        "1",
		name:      "Johaness Cabal",
		role:      "Administrator",
		email:     "jcabal@gmail.com",
		phone:     "052262928",
		contacted: true,
	}
	customers["2"] = Customer{
		id:        "2",
		name:      "Horst Cabal",
		role:      "CEO",
		email:     "hcabal@gmail.com",
		phone:     "0887848118",
		contacted: false,
	}
	customers["3"] = Customer{
		id:        "3",
		name:      "Max Payne",
		role:      "Head of Security",
		email:     "mpayne@gmail.com",
		phone:     "0887848228",
		contacted: true,
	}
}
