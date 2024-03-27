package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter text,thanks: ")

	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	text = strings.Trim(text, "\n")
	fmt.Printf("The text is %v characters long\n", utf8.RuneCountInString(text))

	//sl es un slice
	sl := strings.Split(text, " ")
	//se le añade al slice una entrada más
	sl = append(sl, "!!!")
	//el slice se imprime como string
	fmt.Printf("Salida: %s\n", sl)

}
