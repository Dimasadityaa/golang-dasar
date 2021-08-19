//parameter yang berada di posisi terkahir, memiliki kemampuan dijadikan sebuah varargs
//varargs artinya datanya bisa menerima lebih dari satu input, atatu anggap saja semacam array

//perbedaan parameter biasa dengan tipe data array:
//-jika parameter tipe array, kita wajib membuat array terlebih dahulu sebelum mengirimkan ke function
//-jika parameter menggunakan varargs, kita bisa langsung mengirim datanya, jika lebih dari satu, cukup gunakan tanda koma

package main 

import "fmt" 

func sumAll(number ...int) int {
	total := 0
	for _, value := range number {
		total += value
	}
	return total
}

func main () {
	total := sumAll(10, 10, 10, 10, 10, 10)
	fmt.Println(total)

	slice := []int {10, 20, 20, 20, 20, 20}
	total = sumAll(slice...)
	fmt.Println(total)
}
