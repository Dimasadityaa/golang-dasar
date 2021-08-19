package main 

import "fmt" 

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "esa"
	middleName = "dimas"
	lastName = "aditya"

	return 
}

func main () {
	// fmt.Println(getCompleteName())
	a, b, c := getCompleteName()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}