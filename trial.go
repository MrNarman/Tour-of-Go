package main

import "fmt"

func main() {
	numbers := [6]int{4, 12, 7, 16, 19}

	sum := 0
	count := 0
	for _, num := range numbers {
		if num > 10 {
			sum += num
			count++
		}
	}
	fmt.Println(count)
	fmt.Print(sum)
}
