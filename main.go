package main

import (
	"fmt"

	"github.com/gisanglee/learngo/dayone"
	"github.com/gisanglee/learngo/something"
)

func multiply(a int, b int) int {
	return a * b
}

func main() {
	fmt.Println("Hello World")
	something.SayHello()

	const name string = "gisang"
	var lastName string = "Lee"
	fmt.Println(name)
	fmt.Println(lastName)
	fmt.Println(multiply(2, 2))
	totalLength, upperName := dayone.LenAndUpper(name)
	fmt.Println(totalLength, upperName)
	dayone.RepeatMe("react", "vue", "django", "angular", "next.js", "spring")

	totalLength2, upperName2 := dayone.LenAndUpper2(name)
	fmt.Println(totalLength2, upperName2)

	result := dayone.SuperAdd(1, 2, 3, 4, 5)
	fmt.Println(result)

	teen := dayone.CanISmoke(18)
	adult := dayone.CanISmoke(21)
	fmt.Println(teen, adult)
}
