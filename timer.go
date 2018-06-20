package main

import (
	"fmt"
	"time"
)

func timer() (err error) {
	fmt.Println("Prototype Timer, hit Ctrl-C to stop")
	hideCursor()
	start := time.Now()
	for {
		fmt.Print(time.Now().Sub(start))
		time.Sleep(time.Duration(500000) * time.Nanosecond)
		wipeLine()
	}
	fmt.Println()
	fmt.Println("Your time is:", time.Now().Sub(start))
	showCursor()
	return
}
