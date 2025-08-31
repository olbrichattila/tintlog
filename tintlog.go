package tintlog

import "fmt"

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
)

func Info(msg string) {
	fmt.Println(colorYellow + msg + colorReset)
}

func Error(msg string) {
	fmt.Println(colorRed + msg + colorReset)
}

func Success(msg string) {
	fmt.Println(colorGreen + msg + colorReset)
}
