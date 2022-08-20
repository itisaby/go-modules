package main

type(
	Customer struct {
		FirstName string
		LastName  string
		fullName  string
	}
)

func GetCustomer()(customers []Customer){
		mar := Customer{
			FirstName: "Mary",
			LastName:  "Smith",
			fullName:  "Mary Smith",
			}
		john := Customer{
			FirstName: "John",
			LastName:  "Doe",
			fullName:  "John Doe",
			}
		customers = append(customers, mar)
		customers = append(customers, john)
		return customers
}