package main

import (
	"fmt"
	"time"
)

type Performance struct {
	Date   time.Time     `json:date`
	Chrono time.Duration `json:chrono`
}

func (p Performance) String() string {
	return fmt.Sprintf("▶ %s ◉ %s", p.Date.Format(TIMESTAMP_FORMAT), p.Chrono)
}

func inputPerformance(appfile string) (err error) {
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
	if err = persist(appfile); err != nil {
		return
	}
	showPerformances(10)
	return
}

func appendPerformance(appfile string, p Performance) (err error) {
	s.Performances = append(s.Performances, p)
	return persist(appfile)
}

func showPerformances(limit int) (err error) {
	ln := len(s.Performances)
	if ln == 0 {
		return fmt.Errorf("No performance recorded")
	}
	if limit == 0 {
		limit = ln
	}
	if limit > ln {
		limit = ln
	}
	var total int64
	var n int
	for idx, p := range s.Performances[ln-limit:] {
		n = idx + 1
		total = total + int64(p.Chrono)
		//fmt.Printf("%03d %s\n", n, p)
		fmt.Println(p)
	}
	fmt.Println("Average:", time.Duration(total/int64(n)))
	return
}
