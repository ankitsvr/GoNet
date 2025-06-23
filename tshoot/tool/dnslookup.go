package tool

import (
	"fmt"
	"net"
)

type DnsLookup struct{}

func (d DnsLookup) RuN(target string) error {
	fmt.Printf("Resolving DNS for %s...\n", target)
	ips, err := net.LookupHost(target)
	if err != nil {
		return err
	}
	for _, ip := range ips {
		fmt.Println("IP:", ip)
	}
	return nil
}
