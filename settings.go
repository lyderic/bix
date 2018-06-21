package main

import (
	"os"
	"path/filepath"
)

type Settings struct {
	ruwix64 string
}

var s Settings

func init() {
	s.ruwix64 = filepath.Join(os.Getenv("HOME"), "Downloads", "ruwix.b64")
}
