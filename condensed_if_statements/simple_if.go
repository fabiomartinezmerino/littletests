package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// el sentido de este programa es ejemplificar la sentencia if condensada
// por argumento se le pasa un float que se imprime después de testear que el parseo es correcto
func main() {
	log.Printf("Arrancando...")
	for index, value := range os.Args {
		if value_parsed, err := strconv.ParseFloat(value, 64); err != nil {
			log.Printf("El índice %d no puede parsearse como un float - su valor es %s\n", index, value)
		} else {
			log.Printf("Se puede parsear\n")
			fmt.Printf("En la posición %d está el valor %6.2f\n", index, value_parsed)
		}
	}

}
