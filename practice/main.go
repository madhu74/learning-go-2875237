package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your string: ")
	input, _ := reader.ReadString('\n')
	fmt.Print("Your string is: ", input)
}
