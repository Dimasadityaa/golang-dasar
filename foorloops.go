package main 

import "fmt" 

func main () {
	buah := []string {"pisang", "lemon", "apel"}
	for index, value := range buah {
		fmt.Println(index, value)
	}

	hewan := []string {"kucing", "tikus", "ayam"}
	for i := 0; i < len(hewan); i++ {
		fmt.Println(i, hewan[i])
	}

}