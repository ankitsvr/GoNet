package tool

import (
	"fmt"

	"github.com/go-ping/ping"
)

type PingTool struct{}

func (p PingTool) RuN(target string) error {
	pinger, err := ping.NewPinger(target)
	if err != nil {
		return err
	}
	pinger.Count = 3
	pinger.Run()
	stats := pinger.Statistics()
	fmt.Printf("Ping %s: %v packets, avg delay: %v\n", target, stats.PacketsRecv, stats.AvgRtt)
	return nil
}
