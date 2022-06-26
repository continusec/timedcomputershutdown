package main

import (
	"log"
	"os/exec"
)

func main() {
	err := exec.Command("/sbin/shutdown", "-h", "now").Run()
	if err != nil {
		log.Fatal(err)
	}
}
