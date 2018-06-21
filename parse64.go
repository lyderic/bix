package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func parse64() (err error) {
	var raw []byte
	if raw, err = ioutil.ReadFile(s.ruwix64); err != nil {
		return
	}
	encoded := string(raw)
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
