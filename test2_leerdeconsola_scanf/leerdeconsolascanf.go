package main

import (
	"fmt"
)

func main() {
	var i int
	fmt.Println("Introduce un entero por favor")
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("El número es: %d\n", i)
	}

}
