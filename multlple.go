package main

import "fmt"

func getFullName() (string, string, string) {
	return "eko", "eko", "eko"
}
func main() {
	firstName, _, _ := getFullName()
	fmt.Println(firstName)
	// fmt.Println(midleName)
	// fmt.Println(lastName)
}