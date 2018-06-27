package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Performance struct {
	Date   time.Time     `json:date`
	Chrono time.Duration `json:chrono`
}

func (p Performance) String() string {
	return fmt.Sprintf("▶ %s ◉ %s", p.Date.Format(TIMESTAMP_FORMAT), p.Chrono)
}

func addPerformance(appfile string) (err error) {
	var p Performance
	var dtext, ctext string
	if dtext, err = input("Date (YYYYMMDD HHMMSS)? "); err != nil {
		return
	}
	if len(dtext) == 0 {
		p.Date = time.Now()
	} else {
		if p.Date, err = time.Parse("20060102 150405", dtext); err != nil {
			return
		}
	}
	if ctext, err = input("Chrono (MmS.Cs)? "); err != nil {
		return
	}
	if p.Chrono, err = time.ParseDuration(ctext); err != nil {
		return
	}
	s.Performances = append(s.Performances, p)
	persist(appfile)
	showPerformances()
	return
}

func input(prompt string) (text string, err error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, err = reader.ReadString('\n')
	text = strings.TrimSpace(text)
	return
}

func showPerformances() (err error) {
	for idx, p := range s.Performances {
		fmt.Printf("%03d %s\n", idx+1, p)
	}
	return
}
