package main 

import "fmt" 

func main () {
	var hari = [...]string {
		"senin",
		"selasa",
		"rabu",
		"kamis",
		"jumat",
		"sabtu",
		"minggu",
	}
	fmt.Println(hari)

	var libur [3]int 
	libur[0] = 12
	libur[1] = 13
	libur[2] = 14
	fmt.Println(libur)


}