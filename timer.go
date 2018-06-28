package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func timer() (err error) {
	fmt.Println("Timer - hit [SPACE] to toggle start/stop")
	var terminal Terminal
	if err = terminal.raw(); err != nil {
		return
	}
	start := time.Now()
	// the next two lines to prevent Ctrl-C to be pressed as it messes up the
	// terminal
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	loop()
	fmt.Println("Your time is:", time.Now().Sub(start))
	return terminal.restore()
}

func loop() {
	start := time.Now()
	var b []byte = make([]byte, 1)
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
}
