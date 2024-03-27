package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// se espera que se invoque este programa con una serie de parámetros que deben representar números decimales

	var (
		numberf float64
		err     error
	)

	log.Printf("Empiza la ejecución del programa")
	log.Printf("Nombre del programa: %s\n", os.Args[0])
	for index, value := range os.Args {
		numberf, err = strconv.ParseFloat(value, 32)
		if err != nil {
			log.Printf("En la posición %d no hay float, hay esto: %v\n", index, value)
		} else {
			log.Printf("Hemos encontrado un float")
			fmt.Printf("En la posición %d hay un float, este: %f\n", index, numberf)
		}
	}
}
