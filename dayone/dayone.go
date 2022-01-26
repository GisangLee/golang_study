package dayone

import (
	"fmt"
	"strings"
)

func LenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func LenAndUpper2(name string) (length int, uppercase string) {
	defer fmt.Println("==== 이건 콜백과 비슷한건가 ====")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func RepeatMe(words ...string) {
	fmt.Println(words)
}

func SuperAdd(numbers ...int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}

func CanISmoke(age int) bool {
	if koreanAge := age + 2; koreanAge < 22 {
		return false
	} else {
		return true
	}
}

func CanISmoke2(age int) bool {
	switch {
	case age < 20:
		return false
	case age >= 20:
		return true
	}
}
