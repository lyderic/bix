package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type Settings struct {
	Created      time.Time     `json:created`
	Accessed     time.Time     `json:accessed`
	Performances []Performance `json:performances`
}

var s Settings

func init() {
	var err error
	if _, err = os.Stat(appfile); os.IsNotExist(err) {
		fmt.Println("creating new application file:", appfile)
		s.Created = time.Now()
	} else {
		if err = load(); err != nil {
			log.Fatal("failed to load application file:", err)
		}
	}
	s.Accessed = time.Now()
	if err = persist(); err != nil {
		log.Fatal("failed to write application file:", err)
	}
}

func load() (err error) {
	fmt.Println("LOADING:", appfile)
	var buffer []byte
	if buffer, err = ioutil.ReadFile(appfile); err != nil {
		return
	}
	if err = json.Unmarshal(buffer, &s); err != nil {
		return
	}
	return
}

func persist() (err error) {
	fmt.Println("WRITING:", appfile)
	var buffer []byte
	if buffer, err = json.MarshalIndent(s, "", "  "); err != nil {
		return
	}
	ioutil.WriteFile(appfile, buffer, 0600)
	return
}
