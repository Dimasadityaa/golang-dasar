package main 

import "fmt" 

func main () {
	var golang = 80
	var sql = 80

	var syarat1 = golang >= 80
	var syarat2 = sql >= 80

	var diterima = syarat1 && syarat2

	fmt.Println(diterima)
}