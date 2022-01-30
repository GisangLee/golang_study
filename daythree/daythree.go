package daythree

import "fmt"

func Mypointer(num int) {
	a := num
	b := &a
	*b = num + 29
	fmt.Println(a, &a, b)
}
