package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func timer() (err error) {
	fmt.Println("Timer - hit [SPACE] to stop")
	if err = setTerminal(); err != nil {
		return
	}
	start := time.Now()
	// the next two lines to prevent Ctrl-C to be pressed as it messes up the
	// terminal
	c := make(chan os.Signal)
	//signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	signal.Notify(c, os.Interrupt)
	loop()
	fmt.Println("Your time is:", time.Now().Sub(start))
	return restoreTerminal()
}

func setTerminal() (err error) {
	hideCursor()
	return stty("-icanon", "-echo", "min", "0", "time", "0")
}

func restoreTerminal() (err error) {
	showCursor()
	return stty("sane")
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
