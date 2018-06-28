package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func notyet() (err error) {
	fmt.Println("Not implemented yet")
	return nil
}

func input(prompt string) (text string, err error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, err = reader.ReadString('\n')
	text = strings.TrimSpace(text)
	return
}
