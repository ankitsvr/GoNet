package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usages : Gonet <hostname or IP>")
		return
	}

	target := os.Args[1]

	ip := target

	if net.ParseIP(target) = nil {
		resolvedIP := tools.ResolveDNS(target)
		if err !=nil {
			fmt.Prinln("Dns Lookup error")
			return
		}
		fmt.Printf("resolve ip %s", resolvedIP)
	}

}
