package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func parse64() (err error) {
	//encoded := s.Ruwix64
	encoded := ""
	if len(encoded) == 0 {
		return fmt.Errorf("no ruwix base64 data found")
	}
	var decoded []byte
	if decoded, err = base64.StdEncoding.DecodeString(encoded); err != nil {
		return
	}
	var scanner *bufio.Scanner
	scanner = bufio.NewScanner(bytes.NewReader(decoded))
	scanner.Scan()
	firstline := strings.Replace(scanner.Text(), "]", "", 1)
	var bits = strings.Split(firstline, ",")[1:]
	for _, stime := range bits {
		var i int
		if i, err = strconv.Atoi(stime); err != nil {
			return
		}
		fmt.Println(time.Duration(int64(i) * 1000 * 1000))
	}
	return
}
