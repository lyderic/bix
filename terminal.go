package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

type Terminal struct {
	state string
}

func (terminal *Terminal) init() (err error) {
	var buffer []byte
	cmd := exec.Command("stty", "--save")
	cmd.Stdin = os.Stdin
	if buffer, err = cmd.Output(); err != nil {
		log.Println("cannot save terminal state:", err)
		return
	}
	// the last character of output of stty --save is a '\n'!
	terminal.state = string(buffer[:len(buffer)-1])
	return
}

func (terminal *Terminal) raw(min, intr string) (err error) {
	fmt.Print("\033[?25l") // hide cursor
	err = stty("intr", intr, "-icanon", "-echo", "min", min, "time", "0")
	return
}

func (terminal *Terminal) restore() (err error) {
	fmt.Print("\033[?25h") // show cursor
	return stty(terminal.state)
}

func stty(args ...string) (err error) {
	cmd := exec.Command("stty", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	cmd.Start()
	return cmd.Wait()
}
