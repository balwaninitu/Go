package main

import "fmt"

func main() {

	text := "The following is the account information."

	firstName := "Luke"

	lastName := "Skywalkter"

	age := 20

	weight := 73.0

	height := 1.72

	remainingCredits := 123.55

	accountName := "admin"

	password := "password"

	suscribeUser := true

	fmt.Printf("Values: %s  Type: %T\n", text, text)

	fmt.Printf("Values: %s  Type: %T\n", firstName, firstName)

	fmt.Printf("Values: %s  Type: %T\n", lastName, lastName)

	fmt.Printf("Values: %d  Type: %T\n", age, age)

	fmt.Printf("Values: %.1f  Type: %T\n", weight, weight)

	fmt.Printf("Values: %.2f  Type: %T\n", height, height)

	fmt.Printf("Values: %.2f  Type: %T\n", remainingCredits, remainingCredits)

	fmt.Printf("Values: %s  Type: %T\n", accountName, accountName)

	fmt.Printf("Values: %s  Type: %T\n", password, password)

	fmt.Printf("Values: %t  Type: %T\n", suscribeUser, suscribeUser)

}
