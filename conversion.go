package main 

import "fmt" 

func main () {
	var nama = "dimas"
	var e = nama[0]
	var eString = string(e)
	fmt.Println(nama) //output: dimas
	fmt.Println(eString) //output: d

	var nilai32 int32 = 10000
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)
	fmt.Println(nilai8)
	fmt.Println(nilai64)
}