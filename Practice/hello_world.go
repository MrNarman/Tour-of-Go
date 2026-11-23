package practice

import "fmt"

func HelloWorld() {
	fmt.Println("Hello world")

	//Comments
	/*
		multi line comments
	*/

	var firstName string = "Godrick"
	secondName := "Narman"
	fmt.Printf("Welcome %v %v \n", firstName, secondName)

	var (
		age           int    = 23
		favouriteWord string = "Niaje"
		country       string = "kenya"
	)

	fmt.Printf("Details: \n\t %v \n\t %v \n\t %v \n", age, favouriteWord, country)

	const PI = 3.142
	const SHORTERPI float32 = 3.1
	fmt.Printf("mathematical PI: %v \n", PI)
	fmt.Printf("shorter PI: %v \n", SHORTERPI)

	var a, b string = "mjini", "nairobi"
	fmt.Print(a, "\n", b, " ni kuzuri mno ", "\n")

	fmt.Printf("Go format %#v", a)
	fmt.Printf("\n Normal format %v", b)
}
