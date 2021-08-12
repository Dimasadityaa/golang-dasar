package main 

import "fmt" 

func main () {
	var nama = ""

	switch nama {
	case "dimas":
		fmt.Println("dimas")
	case "dwi":
		fmt.Println("dwi")
	default:
		fmt.Println("input name")
	}
}