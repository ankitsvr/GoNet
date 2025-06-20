package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("please enter the url/hostname")
		os.Exit(1)
	}
	target := os.Args[1]

	ip := net.ParseIP(target)
	if ip != nil {
		names, err := net.LookupIP(target)
		if err != nil {
			fmt.Println("failed to do reverse lookup")
			return

		}
		fmt.Printf("hostname for the give ip is : %s ", names)

	} else {
		IP, err := net.LookupHost(target)
		if err != nil {
			fmt.Println("DNS lookup failed")

		}
		fmt.Printf("IP for the given host %s: %v", target, IP)
	}
}
