package main

import (
	"fmt"
	"os"
	"time"
)

func timer() (err error) {
	fmt.Println("Prototype Timer, hit [SPACE] to stop")
	stty("-icanon", "-echo", "min", "0", "time", "0")
	defer stty("sane")
	start := time.Now()
	var b []byte = make([]byte, 1)
	hideCursor()
	defer showCursor()
	for {
		time.Sleep(time.Duration(500000) * time.Nanosecond)
		fmt.Printf("\r%s\r", time.Now().Sub(start))
		os.Stdin.Read(b)
		if b[0] == 10 { // enter
			continue
		}
		if b[0] == 32 { // space
			break
		}
	}
	fmt.Println("Your time is:", time.Now().Sub(start))
	return
}
