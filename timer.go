package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func timer() (err error) {
	fmt.Println("Timer - hit [SPACE] to toggle start/stop")
	fmt.Print("0m00s")
	var terminal Terminal
	if err = terminal.init(); err != nil {
		return
	}
	if err = spacePressed(terminal); err != nil {
		return
	}
	start := time.Now()
	// the next two lines to prevent Ctrl-C to be pressed as it messes up the
	// terminal
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	if err = loop(terminal); err != nil {
		  return
	}
	fmt.Println("Your time is:", time.Now().Sub(start))
	return terminal.restore()
}

func spacePressed(terminal Terminal) (err error) {
	if err = terminal.raw("1"); err != nil {
		return
	}
	var b []byte = make([]byte, 1)
	os.Stdin.Read(b)
	if b[0] != 32 {
		return fmt.Errorf("Please press [SPACE] to start the timer!")
	}
	return
}

func loop(terminal Terminal) (err error) {
	if err = terminal.raw("0"); err != nil {
		return
	}
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
	return
}
