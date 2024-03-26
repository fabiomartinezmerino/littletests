package main

import (
	"fmt"
)

func main() {
	var i int
	_, err := fmt.Scanf("%d", &i)
	fmt.Printf("El número es: %d\n", i)
	_ = err
}
