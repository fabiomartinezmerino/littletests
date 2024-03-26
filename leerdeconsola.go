package main

import "fmt"

func main() {

	var i int
	_, err := fmt.Scanf("Input a number %d", &i)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("El número introducido es: %d", i)
	}

}
