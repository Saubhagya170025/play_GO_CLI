package ping

import (
	"fmt"
	"net"
	"time"
)

func Run(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: godevops ping <host>")
		return
	}

	host := args[0]

	start := time.Now()
	_, err := net.DialTimeout("tcp", host+":80", 3*time.Second)
	elapsed := time.Since(start)

	if err != nil {
		fmt.Println("Host unreachable:", err)
		return
	}

	fmt.Printf("Host %s reachable in %v\n", host, elapsed)
}