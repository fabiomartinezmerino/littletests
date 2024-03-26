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
	fmt.Println("Enter text: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	text = strings.Trim(text, "\n")
	fmt.Printf("The text is %v characters long\n", utf8.RuneCountInString(text))

	sl := strings.Split(text, " ")

	sl = append(sl, "!!!")
	output := strings.Join(sl, " ")
	fmt.Printf("Salida: %s\n", output)

}
