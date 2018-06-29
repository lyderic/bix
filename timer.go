package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

var chrono time.Duration

func timer(appfile string) (err error) {
	fmt.Println("Timer - hit [SPACE] to toggle start/stop")
	fmt.Print("0m00s")
	var terminal Terminal
	if err = terminal.init(); err != nil {
		return
	}
	var ctx context.Context
	var cancel context.CancelFunc
	ctx, cancel = context.WithCancel(context.Background())
	done := make(chan struct{})
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go stop(terminal, c, cancel)
	if err = spacePressed(terminal); err != nil {
		return
	}
	go loop(terminal, ctx, done)
	<-done
	fmt.Println("Your time is:", chrono)
	return record(appfile)
}

func stop(terminal Terminal, c chan os.Signal, cancel context.CancelFunc) (err error) {
	select {
	case <-c:
		if err = terminal.restore(); err != nil {
			return
		}
		cancel()
	}
	return
}

func loop(terminal Terminal, ctx context.Context, done chan struct{}) (err error) {
	if err = terminal.raw("0", " "); err != nil {
		return
	}
	start := time.Now()
	for {
		time.Sleep(time.Duration(750000) * (time.Nanosecond))
		chrono = time.Now().Sub(start)
		select {
		case <-ctx.Done():
			done <- struct{}{}
		default:
			fmt.Printf("\r%v\r", chrono)
		}
	}
	return
}

func record(appfile string) (err error) {
	var answer string
	if answer, err = input("Record time [y/N]? "); err != nil {
		return
	}
	if len(answer) == 0 || answer[0] == 'n' || answer[0] == 'N' {
		fmt.Println("Time not recorded")
		return
	}
	var p Performance
	p.Date = time.Now()
	p.Chrono = chrono
	appendPerformance(appfile, p)
	showPerformances(10)
	return
}

func spacePressed(terminal Terminal) (err error) {
	if err = terminal.raw("1", ""); err != nil {
		return
	}
	var b []byte = make([]byte, 1)
	os.Stdin.Read(b)
	if b[0] != 32 {
			if err = terminal.restore(); err != nil {
				  return
			}
			return fmt.Errorf("\rPlease press [SPACE] to start the timer!")
	}
	return terminal.restore()
}
