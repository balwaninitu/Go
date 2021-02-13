package main

import (
	"fmt"
)

func main() {
	var userShoppingListMenuInput int
	fmt.Print("Enetr your choice:")
	fmt.Scanln(&userShoppingListMenuInput)

	if userShoppingListMenuInput < 1 || userShoppingListMenuInput > 5 {
		fmt.Println("Enter valid choice between 1 to 5")
	}

	// validInput := regexp.MustCompile(`^[1-6]*$`)
	// isNum := validInput.Match([]byte(userShoppingListMenuInput))
	// fmt.Print(isNum)

	// if len(strings.TrimSpace(userShoppingListMenuInput)) == 0 {
	// 	fmt.Println("\nString is empty!")
	// }

}
