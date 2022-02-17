package main

import "fmt"

//main used to run all other function
func main() {
	//fmt.Println(checkTriangle(1, 10, 1))
	for _, value := range pascalTriangle(8) {
		fmt.Printf("%d\n", value)
	}
}
