package main

import (
	"fmt"
	"github.com/go-ping/ping"
	"os"
	"runtime"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s <ip> <timeout seconds>\n", os.Args[0])
		os.Exit(1)
	}

	host := os.Args[1]
	giveupTimeout, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Invalid timeout value: %s\n", os.Args[2])
		os.Exit(1)
	}

	pingTimeout := 3 * time.Second

	exitCode := 1
	startTime := time.Now()

	for time.Since(startTime).Seconds() < float64(giveupTimeout) {
		pinger, err := ping.NewPinger(host)
		if err != nil {
			fmt.Printf("ERROR: %s\n", err.Error())
			os.Exit(1)
		}

		if runtime.GOOS == "linux" {
			pinger.SetPrivileged(true)
		}

		pinger.Count = 1
		pinger.Timeout = pingTimeout

		pinger.OnRecv = func(pkt *ping.Packet) {
			exitCode = 0
		}

		pinger.Run()

		// If a packet was received, exit the program
		if exitCode == 0 {
			break
		}

		// Sleep before retrying
		time.Sleep(1 * time.Second)
	}

	os.Exit(exitCode)
}
