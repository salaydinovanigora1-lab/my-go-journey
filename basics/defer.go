package main

import "fmt"

func main() {

	process(10)
}

func process(i int) {
	defer fmt.Println("Deferred i value:", i)
	defer fmt.Println("first Deferred statement executed")
	defer fmt.Println("second Deferred statement executed")
	defer fmt.Println("third Deferred statement executed")
	i++
	fmt.Println("Normal execution statement")
	fmt.Println("Value of i:", i)
}
