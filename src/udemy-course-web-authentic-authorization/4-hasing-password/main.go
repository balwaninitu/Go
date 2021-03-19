package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	pass := "123456789"

	hashedPass, err := hashpassword(pass)
	if err != nil {
		panic(err)
	}

	err = comparePassword(pass, hashedPass)
	if err != nil {
		log.Fatalln("Not logged in")
	}
	log.Println("Logged in!")

	fmt.Println("Password string:", pass)
	fmt.Println("Hashed Password []byte :", hashedPass)
}

//takes password string and return slice of byte
//default cost is 10
//there are 3 cost to choose min cost is 4 and max is 31
//more cost will take longer time so default is genarally chosen
//one advantage of this method is user can put anything in string it can be emoji as well
func hashpassword(password string) ([]byte, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("Error while generating bcrypt hash from password: %w", err)
	}
	return bs, nil
}
func comparePassword(password string, hashedPass []byte) error {
	//hashpassword func has return []byte hence password string convert to []byte
	err := bcrypt.CompareHashAndPassword(hashedPass, []byte(password))
	if err != nil {
		return fmt.Errorf("Invalid password: %w", err)
	}
	return nil
}
