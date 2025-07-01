package main

import (
	"GoNet/tasks"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 || os.Args[1] != "check" {
		fmt.Println("Usage: go run main.go check <hostname>")
		return
	}

	host := os.Args[2]
	fmt.Println("Running checks for:", host)

	dnsCh := make(chan string)
	pingCh := make(chan string)
	traceCh := make(chan string)
	mtrCh := make(chan string)

	go tasks.PerformDNSLookup(host, dnsCh)
	go tasks.RunPing(host, pingCh)
	go tasks.RunTraceroute(host, traceCh)
	go tasks.RunMTR(host, mtrCh)

	fmt.Println("\n--- DNS Lookup ---")
	fmt.Print(<-dnsCh)

	fmt.Println("\n--- Ping ---")
	fmt.Print(<-pingCh)

	fmt.Println("\n--- Traceroute ---")
	fmt.Print(<-traceCh)

	fmt.Println("\n--- MTR ---")
	fmt.Print(<-mtrCh)
}
