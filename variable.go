package main

import "fmt"

func main () {
	var nama string = "uus"
	nama = "dimas"
	fmt.Println(nama) //output dimas
	//nilai variabel bisa dibuah

	// variabel multi

	var (
		firstName = "dimas"
		lastName = "aditya"
	)
	fmt.Println(firstName, lastName) //output: dimas aditya
}