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

    ports := [20]int{21, 22, 23, 25, 53, 80, 110, 111, 135, 139, 143, 443, 445, 993, 995, 1723, 3306, 3389, 5900, 8080}


	var results []scan_result

	for _, s := range ports {
		results = append(results, scan_port(hostname, s))
	}

	return results
}

func main() {
	var host string
	fmt.Println("Port Scanner by mmkamron")
	fmt.Printf("Enter an ip address or url to scan: ")
	fmt.Scan(&host)
	op := initial_scan(host)
    for _, r := range op {
        if (r.State == "Open") {
            fmt.Println(r)
        }
    }
}
