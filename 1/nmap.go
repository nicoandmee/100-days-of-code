package main

import (
	"fmt"
	"os"
	"os/exec"
	"q"
	"syscall"
)

func main() {
	// usage: go run nmap.go [ip]

	// check the program was supplied proper # of args
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip\n", os.Args[0])
	}

	bin, err := exec.LookPath("/usr/bin/nmap")
	if err != nil {
		q.Q(err) // DEBUG
	}

	if len(bin) == 0 {
		fmt.Println("Could not locate nmap binary on system, exiting.")
		os.Exit(1)
	}

	// create nmap args
	args := []string{"nmap", "-sV", "-A", os.Args[1]}
	env := os.Environ()

	execErr := syscall.Exec(bin, args, env)
	if execErr != nil {
		q.Q(execErr) // DEBUG
	}

}
