package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func timer() (err error) {
	fmt.Println("Prototype Timer, Ctrl-C to stop")
	exec.Command("stty", "-F", "/dev/tty", "-icanon", "-echo", "min", "1", "time", "10").Run()
	defer exec.Command("stty", "-F", "/dev/tty", "sane").Run()
	start := time.Now()
	for {
		time.Sleep(time.Duration(500000) * time.Nanosecond)
		fmt.Printf("\r%s\r", time.Now().Sub(start))
		var b []byte
		os.Stdin.Read(b)
		if len(b) == 0 { continue }
		if b[0] == 32 {
			break
		}
	}
	fmt.Println()
	fmt.Println("Your time is:", time.Now().Sub(start))
	return
}
