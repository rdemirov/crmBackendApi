package customers

type Customer struct {
	id        string
	name      string
	role      string
	email     string
	phone     string
	contacted bool
}

var customers []Customer

func initializeCustomersDb() {
	customers = append(customers, Customer{
		id:        "1",
		name:      "Johaness Cabal",
		role:      "Necromancer",
		email:     "jcabal@gmail.com",
		phone:     "0895453490",
		contacted: true,
	})
	customers = append(customers, Customer{
		id:        "2",
		name:      "Max Payne",
		role:      "Policeman",
		email:     "mpayne@gmail.com",
		phone:     "08775453490",
		contacted: false,
	})
	customers = append(customers, Customer{
		id:        "3",
		name:      "Horst Cabal",
		role:      "Vampire",
		email:     "hcabal@gmail.com",
		phone:     "08885353490",
		contacted: true,
	})
}
