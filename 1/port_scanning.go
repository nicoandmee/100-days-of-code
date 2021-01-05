package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"q"
	"strconv"
	"time"
)

var targetIP = "173.233.75.183"
var minPort = 1
var maxPort = 1024

//TODO: Add support for command line args, custom threading count

func main() {
	fmt.Println("I'm a port scanner!")

	// check we have the right # of args
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s -ip [target]\n", os.Args[0])
		os.Exit(1)
	}

	target := flag.String("ip", "localhost", "IP Address to scan")
	flag.Parse()
	fmt.Printf("Scanning: %s\n", *target)

	activeThreads := 0
	doneChannel := make(chan bool)

	for port := minPort; port <= maxPort; port++ {
		go connect(*target, port, doneChannel) // <-- go threading/routine
		activeThreads++
	}

	// wait for all threads to finish
	for activeThreads > 0 {
		<-doneChannel
		activeThreads--
	}
}

func connect(ip string, port int, doneChannel chan bool) {
	_, err := net.DialTimeout("tcp", ip+":"+strconv.Itoa(port), time.Second*10)
	if err == nil {
		log.Printf("Host %s | Found open port: %d\n", ip, port)
	}
	if err != nil {
		q.Q(err) // DEBUG
	}
	doneChannel <- true
}
