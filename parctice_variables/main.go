package main

import "fmt"

const variable1 = 100

func main() {

	var anInt int = 23
	fmt.Printf("This is of the type %T\n", anInt)
	fmt.Println("Vale is", anInt)
	var aString string = "This is my string"
	fmt.Printf("This is of the type %T\n", aString)
	fmt.Println("Value is", aString)
	var anInt2 int
	anInt2 = 30
	fmt.Printf("This is of the type %T\n", anInt2)
	fmt.Println("Value is", anInt2)
	var aString2 = "What a String" // implicit declaration
	fmt.Printf("This is of the type %T\n", aString2)
	fmt.Println("Value is:", aString2)
	aString3 := "What a String2" // implicit declaration type not mentioned
	fmt.Printf("This is of the type %T\n", aString3)
	fmt.Println("Value is:", aString3)
	// implicit declaration type not mentioned
	aString4 := "What a String3"
	fmt.Printf("This is of the type %T\n", aString4)
	fmt.Println("Value is:", aString4)
	fmt.Println("constant", variable1)
	aString4 = "now what"
	fmt.Println("Value is:", aString4)
}
