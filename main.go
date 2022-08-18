package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/common-nighthawk/go-figure"
)

type Result struct {
	Port  int
	State string
	Name  string
}

func scan_port(hostname string, port int) Result {
	result := Result{Port: port}
	address := net.JoinHostPort(hostname, strconv.Itoa(port))
	conn, err := net.DialTimeout("tcp", address, 1*time.Second)
	if err != nil {
		result.State = "Closed or filtered"
		return result
	}
	defer conn.Close()
	// result.State = "Open"
	result.State = "Open"
	return result
}

func initial_scan(hostname string) []Result {
	var results []Result

	for k, v := range knownPorts {
		res := scan_port(hostname, k)
		res.Name = v
		results = append(results, res)
	}

	return results
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func main() {
	addr := flag.String("u", "localhost", "IP address/Domain name")
	flag.Parse()
	myFigure := figure.NewFigure("Goscan", "", true)
	myFigure.Print()
	for _, r := range initial_scan(*addr) {
		if r.State == "Open" {
			fmt.Printf("Port %d is open | %s\n", r.Port, r.Name)
		}
	}
}
