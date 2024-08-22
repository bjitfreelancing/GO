package main

import "fmt"

func main() {
	// fmt.Println("Hello World!")

	// var name string = "Soumyadeep"
	// fmt.Printf("Hi %s, have a nice day\n", name)
	// name1 := "Soumyadeep"
	// fmt.Printf("Hi %s, have a nice day\n", name1)

	// var (
	// 	x, y   int    = 100, 200
	// 	result int    = x + y
	// 	str    string = "sum of %d and %d: %d\n"
	// )
	// fmt.Printf(str, x, y, result)
	var c1 complex128 = 2 + 3i
	var c2 complex128 = 5 + 7i
	c := c1 + c2
	fmt.Println(c)

	const pi = 3.14
	fmt.Println(pi)

	const (
		red    = iota
		green  = iota
		yellow = iota
	)
	fmt.Printf("red = %d, green = %d, yellow = %d\n", red, green, yellow)

	const (
		VERBOSE = 1 << iota
		CONFIG_FROM_DISK
		DATABASE_REQUIRED
		LOGGER_ACTIVATED
		DEBUG
		FLOAT_SUPPORT
		RECOVERY_MODE
		REBOOT_ON_FAILURE
	)
	fmt.Printf("%b, %b, %b, %b, %b, %b, %b, %b\n", VERBOSE, CONFIG_FROM_DISK, DATABASE_REQUIRED, LOGGER_ACTIVATED, DEBUG, FLOAT_SUPPORT, RECOVERY_MODE, REBOOT_ON_FAILURE)
}
