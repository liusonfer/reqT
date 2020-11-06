package utils

import (
	"bufio"
	"fmt"
	"os"
)

func Input(prompt string) string {
	// var text string
	fmt.Println(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, _, _ := reader.ReadLine()
	// fmt.Scanf("%s\n", &text)
	return string(text)
}
