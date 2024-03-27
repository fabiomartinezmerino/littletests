package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

//leemos de línea de comandos una serie de números de exactamente 4 cifras

func main() {
	//create reader to read from commandline
	myreader := bufio.NewReader(os.Stdin)

	//Prompt with instructions
	fmt.Println("Por favor introduce un número de 4 dígitos, después pulsa Enter")

	//Leemos
	numbers, err := myreader.ReadString('\n')

	if err != nil {
		log.Fatal("%v", err)
	}

	numbers = strings.Trim(numbers, "\n") //se elimina saltos de línea

	slcnumbers := strings.Split(numbers, " ") // se colocan los diferentes números en un slice

	for index, value := range slcnumbers {

		//si hay más de 4 runas en cada entrada del slice.... algo está mal formateado
		if utf8.RuneCountInString(value) > 4 {
			log.Fatal(`Bad format`)
		}

		fmt.Printf("La cadena introducida en el lugar %d es: %s\n", index, value)

		number, err := strconv.Atoi(value)

		if err != nil {
			log.Fatal("%v", err)
		}

		fmt.Printf("El entero es: %v\n", number)

	}

}
