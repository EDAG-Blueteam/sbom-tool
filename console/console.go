package console

import (
	"fmt"
	"strings"
)

func Clear() {

	// clear screen and reset cursor
	fmt.Print("\u001b[2J\u001b[0f")

	// clear scroll buffer
	fmt.Print("\u001b[3J")

}

func Error(message string) {

	if strings.Contains(message, "\n") {

		var lines = strings.Split(message, "\n")

		for l := 0; l < len(lines); l++ {
			fmt.Println("\u001b[41m " + lines[l] + "\u001b[K")
		}

		fmt.Println("\u001b[0m")

	} else {
		fmt.Println("\u001b[41m " + message + "\u001b[K\u001b[0m")
	}

}

func Info(message string) {

	if strings.Contains(message, "\n") {

		var lines = strings.Split(message, "\n")

		for l := 0; l < len(lines); l++ {
			fmt.Println("\u001b[42m " + lines[l] + "\u001b[K")
		}

		fmt.Println("\u001b[0m")

	} else {
		fmt.Println("\u001b[42m " + message + "\u001b[K\u001b[0m")
	}

}

func Log(message string) {

	if strings.Contains(message, "\n") {

		var lines = strings.Split(message, "\n")

		for l := 0; l < len(lines); l++ {
			fmt.Println("\u001b[40m " + lines[l] + "\u001b[K")
		}

		fmt.Println("\u001b[0m")

	} else {
		fmt.Println("\u001b[40m " + message + "\u001b[K\u001b[0m")
	}

}

func Warn(message string) {

	if strings.Contains(message, "\n") {

		var lines = strings.Split(message, "\n")

		for l := 0; l < len(lines); l++ {
			fmt.Println("\u001b[43m " + lines[l] + "\u001b[K")
		}

		fmt.Println("\u001b[0m")

	} else {
		fmt.Println("\u001b[43m " + message + "\u001b[K\u001b[0m")
	}

}
