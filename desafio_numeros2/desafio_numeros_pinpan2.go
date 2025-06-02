package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		var result string

		if i%3 == 0 {
			result += "Pin"
		}
		if i%5 == 0 {
			result += " Pan"
		}

		switch result {
		case "":
			fmt.Println(i)
		default:
			fmt.Println(result)
		}
	}
}
