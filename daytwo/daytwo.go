package daytwo

import (
	"fmt"
	"strings"
)

func SwitchCase(name string) (bool, string) {
	switch myName := strings.ToUpper(name); myName {
	case "GISANG":
		return true, "You'r name is GISANG"
	}
	return false, "You'r name is not GISANG"
}

func afterGotoWork(message string) string {
	switch message {
	case "일찍 일어나셨군요.":
		return "버스 탈듯"
	case "회사가 가까우신가요":
		return "택시 탈듯?"
	}
	return "지각할듯?"
}

func GoToWork(hour int) string {
	message := ""

	if hour <= 6 {
		message = "일찍 일어나셨군요."
	} else if hour < 8 {
		message = "회사가 가까우신가요"
	} else {
		message = "지각하는 거 아닌가 몰라"
	}

	defer fmt.Println(afterGotoWork(message))

	return message
}

func Netflix(movie string) string {
	message := strings.ToUpper(movie)

	if message == "NASA" {
		return "this is NASA"
	} else {
		return "this is not NASA"
	}
}
