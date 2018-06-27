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

func setup(appfile string) (err error) {
	if _, err = os.Stat(appfile); os.IsNotExist(err) {
		fmt.Println("creating new application file:", appfile)
		s.Created = time.Now()
	} else {
		if err = load(appfile); err != nil {
			log.Fatal("failed to load application file:", err)
		}
	}
	s.Accessed = time.Now()
	if err = persist(appfile); err != nil {
		log.Fatal("failed to write application file:", err)
	}
	return
}

func load(appfile string) (err error) {
	var buffer []byte
	if buffer, err = ioutil.ReadFile(appfile); err != nil {
		return
	}
	if err = json.Unmarshal(buffer, &s); err != nil {
		return
	}
	return
}

func persist(appfile string) (err error) {
	var buffer []byte
	if buffer, err = json.MarshalIndent(s, "", "  "); err != nil {
		return
	}
	ioutil.WriteFile(appfile, buffer, 0600)
	return
}
