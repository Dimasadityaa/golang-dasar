package main

import "fmt"

func main () {
	var inte8 int8 = 127 //limit 0-127
	var inte16 int16 = 32767 // limit 0-32767
	var inte32 int32 = 2147483647 // limit 0-2147483647
	var inte64 int64 = 9223372036854775807 // limit 0 - 9223372036854775807
	var integer int = 9223372036854775807 // sama dengan int 32 dan 64

	var nol = 0

	fmt.Println("int8:", nol, "-", inte8)
	fmt.Println("int16:", nol, "-", inte16)
	fmt.Println("int32:", nol, "-", inte32)
	fmt.Println("int64:", nol, "-", inte64)
	fmt.Println("integer", nol, "-",integer)
}

// TipeData	Cakupan Nilai
// uint8 	  0 – 255
// uint16	  0 – 65535
// uint32	  0 – 4294967295
// uint64	  0 – 18446744073709551615
// uint	    Sama dengan uint34 dan uint64
// byte	    Sama dengan uint8
// int8	    -128 – 127
// init16	  -32768 – 32767
// int32	    -2147483648 – 2147483647
// int64	    -9223372036854775808 – 9223372036854775807
// int	      Sama dengan int 32 dan 64
// rune	    Sama dengan int32
