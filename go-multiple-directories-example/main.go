package main

import (
	"fmt"
	"C:\Projects\Go\go-multiple-directories-example\models"
	"C:\Projects\Go\go-multiple-directories-example\routes"

)

func main()  {

	fmt.println("Main package - main file")
	models.AllUsers()
	routes.APIPostRoute()
	
}

