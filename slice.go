package main 

import "fmt" 

func main () {
	slice := []string {"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}

	slice2 := slice[5:]

	ap := append(slice2, "seninlagi")
	fmt.Println(ap)

	newslice := make([]string, 2, 4)
	newslice[0] = "golang"
	newslice[1] = "sql"

	fmt.Println(newslice)

	copySlice := make([]string, len(newslice), cap(newslice))
	copy(copySlice, newslice)
	fmt.Println(copySlice)
}