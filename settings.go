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
		var answer string
		if answer, err = input(fmt.Sprintf("Create new application file (%s) [y/N]? ", appfile)); err != nil {
			return
		}
		if len(answer) == 0 || answer[0] == 'n' || answer[0] == 'N' {
			return fmt.Errorf("no application file: %s", appfile)
		}
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
