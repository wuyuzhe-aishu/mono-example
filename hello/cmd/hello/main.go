package main

import (
	"encoding/json"
	"os"
	"runtime/debug"
)

func main() {
	if info, ok := debug.ReadBuildInfo(); ok {
		json.NewEncoder(os.Stdout).Encode(info)
	}
}
