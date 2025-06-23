package main

import (
	"fmt"

	"github.com/ankitsvr/GoNet/tshoot/tool"
)

func Execute(t tool.NetTool, target string) {
	fmt.Println("------ Running Tool ------")
	if err := t.RuN(target); err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("--------------------------\n")
}

func main() {
	target := "www.google.com"
	tools := []tool.NetTool{tool.DnsLookup{}, tool.PingTool{}}

	for _, t := range tools {
		Execute(t, target)
	}
}
