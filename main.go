package main

import (
	"fmt"
	"strconv"
	"net"
	"time"
)

type scan_result struct {
	Port int
	State string
}

func scan_port(hostname string, port int) scan_result {
	result := scan_result{Port: port}
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout("tcp", address, 3*time.Second)
	
	if err != nil {
		result.State = "Closed or filtered"
		return result
	}
	defer conn.Close()
	result.State = "Open"
	return result
}

func initial_scan(hostname string) []scan_result {

	var results []scan_result

	for i := 0; i <= 1024; i++ {
		results = append(results, scan_port(hostname, i))
	}

	return results
}

func main() {
	var host string
	fmt.Println("Port scanner by mmkamron")
	fmt.Println("Enter an ip address or url to scan: ")
	fmt.Scanln(&host)
	op := initial_scan(host)
	fmt.Println(op)
}
