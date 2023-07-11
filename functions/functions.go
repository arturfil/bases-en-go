package functions

import (
	"fmt"
)

func PrintUser() {
    nombre, email, userName, age := getUserDetails()

    fmt.Println("Your user first name is:", nombre) 
    fmt.Println("Your email is:", email) 
    fmt.Println("Your username is:", userName) 
    fmt.Println("Your age is:", age) 
}

func getUserDetails() (string, string, string, int) {
   var name string
   var email string
   var username string
   var age int

   fmt.Println("Please enter your name: ")
   fmt.Scan(&name)

   fmt.Println("Please enter your email: ")
   fmt.Scan(&email)

   fmt.Println("Please enter your username: ")
   fmt.Scan(&username)
   
   fmt.Println("Please enter your age: ")
   fmt.Scan(&age)

   return name, email, username, age
}
