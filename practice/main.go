package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("You entered:", input)
	fmt.Print("Enter a number: ")
	inputNumer, _ := reader.ReadString('\n')
	aNumber, err := strconv.ParseFloat(strings.TrimSpace(inputNumer), 64)
	if err != nil {
		fmt.Println("Encountered an error: ", err)
	} else {
		fmt.Println("The number is : ", aNumber)
	}

}
