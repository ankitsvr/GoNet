package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"

	"github.com/go-ping/ping"
)

//logic to check if the given target is IP or hostname

func isIP(target string) bool {
	return net.ParseIP(target) != nil
}

// logic to ressolve the hostname
func dnsLookup(address string) ([]string, error) {
	return net.LookupHost(address)
}

// logic for the ping
func doPing(target string) {
	fmt.Println("Pinging : ", target)
	pinger, err := ping.NewPinger(target)
	if err != nil {
		fmt.Println("ping error: ", err)
		return
	}
	pinger.Count = 4
	pinger.Run()
	stats := pinger.Statistics()
	fmt.Printf("Results: %+v\n", stats)

}

// logic for the traceroute
func doTrace(target string) {
	fmt.Println("Traceroute to : ", target)
	var cmd *exec.Cmd
	cmd = exec.Command("traceroute", target)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("traceroute Failed: ", err)
		return
	}
	fmt.Println(string(output))
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("usage: go run main <target> 'ip or hostname'")
		return
	}
	input := strings.TrimSpace(os.Args[1])
	target := input

	if !isIP(target) {
		fmt.Println("detected Hostname, Resolving name...\n")
		ips, err := dnsLookup(target)
		if err != nil {
			fmt.Println("DNS lookup failed")
			return
		}
		fmt.Println("Found IPs", ips)
		target = ips[0]

	} else {
		fmt.Println("detected IP address")

	}
	doPing(target)
	doTrace(target)

}
