package main

import "fmt"

func main () {
	const nama = "dimas" //const nilainya tetap tidak bisa dirubah
	fmt.Println(nama)//output: dimas

	// const multi

	const (
		firstName = "dwi"
		lastName = "reydesta"
	)

	fmt.Println(firstName, lastName)
}