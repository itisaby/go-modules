package main

import "fmt"

func main() {
	customer:=GetCustomer()
	for x, customer := range customer {
		fmt.Println(x, customer.fullName)
	}
}

func getData()(customers []string){
	// if customerId == 1 {
	// 	customer = "John"
	// }else if customerId == 2 {
	// 	customer = "Mary"
	// }else{
	// 	customer = "Unknown"
	// }
	// return customer
	// customer:= "John"
	// customers[0] = customer
	// customers[1] = "Mary"
	// customers = append(customers, "Hello")
		customers = append(customers, "John2")
		customers = append(customers, "John1")
		customers = append(customers, "John")
		customers = append(customers, "Mary")
		customers = append(customers, "Hello")
		for x, customer := range customers {
			customer=customers[x]
			fmt.Println(customer)
		}
	return customers

}