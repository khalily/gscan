package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"regexp"
	"scanfactory"
	"strconv"
	"strings"
	. "types"
)

var ipRange *string = flag.String("ip", "127.0.0.1", `IP range of scan. Example:
		192.168.1.1
		192.168.1.1, 192.168.1.5
		192.168.1.1-192.168.1.100`)

var portRange *string = flag.String("p", "1-1024", `Port range of scan. Example:
		135
		135, 445, 3389
		1-1024`)

var way *string = flag.String("w", "connect", `Way of scan. Expample
		connect
		syn
		fin`)

var help *string = flag.String("h", "help", "help doc")

func parsePort(portRange *string) (ports []Port, err error) {
	if strings.Contains(*portRange, ",") {
		for _, port := range strings.Split(*portRange, ",") {
			port, err := strconv.Atoi(port)
			if err != nil {
				return ports, err
			}
			ports = append(ports, Port(port))
		}
	} else if strings.Contains(*portRange, "-") {
		portst := strings.Split(*portRange, "-")
		if len(portst) != 2 {
			err = errors.New("port don't parse")
			return
		}
		ports = parseTwoPortRange(portst[0], portst[1])
	} else {
		port, err := strconv.Atoi(*portRange)
		if err != nil {
			return ports, err
		}
		ports = append(ports, Port(port))
	}
	return
}

func parseTwoPortRange(portStart, portEnd string) (ports []Port) {
	port1, _ := strconv.Atoi(portStart)
	port2, _ := strconv.Atoi(portEnd)

	for port := port1; port != port2+1; port++ {
		ports = append(ports, Port(port))
	}
	return
}

func parseIp(ipRange *string) (ips []Ip, err error) {
	if strings.Contains(*ipRange, ",") {
		for _, ip := range strings.Split(*ipRange, ",") {
			ips = append(ips, Ip(ip))
		}
	} else if strings.Contains(*ipRange, "-") {
		ipst := strings.Split(*ipRange, "-")
		ips = parseTwoIpRange(ipst[0], ipst[1])
	} else {
		ips = append(ips, Ip(*ipRange))
	}
	return
}

func parseTwoIpRange(ipStart, ipEnd string) (ips []Ip) {
	ip1 := Ip2long(ipStart)
	ip2 := Ip2long(ipEnd)

	for ip := ip1; ip != ip2+1; ip++ {
		s := Long2ip(ip)
		var match = false
		if match, _ = regexp.MatchString("255$", s); match {
			continue
		} else if match, _ = regexp.MatchString("0$", s); match {
			continue
		} else {
			ips = append(ips, Ip(Long2ip(ip)))
		}
	}
	return
}

func Long2ip(ip uint32) string {
	return fmt.Sprintf("%d.%d.%d.%d", ip>>24, ip<<8>>24, ip<<16>>24, ip<<24>>24)
}

func Ip2long(ipstr string) (ip uint32) {
	r := `^(\d{1,3})\.(\d{1,3})\.(\d{1,3})\.(\d{1,3})`
	reg, err := regexp.Compile(r)
	if err != nil {
		return
	}
	ips := reg.FindStringSubmatch(ipstr)
	if ips == nil {
		return
	}

	ip1, _ := strconv.Atoi(ips[1])
	ip2, _ := strconv.Atoi(ips[2])
	ip3, _ := strconv.Atoi(ips[3])
	ip4, _ := strconv.Atoi(ips[4])

	if ip1 > 255 || ip2 > 255 || ip3 > 255 || ip4 > 255 {
		return
	}

	ip += uint32(ip1 * 0x1000000)
	ip += uint32(ip2 * 0x10000)
	ip += uint32(ip3 * 0x100)
	ip += uint32(ip4)

	return
}

func main() {
	flag.Parse()

	if len(os.Args) > 1 && os.Args[1] == "-h" ||
		len(os.Args) == 1 {
		flag.Usage()
		os.Exit(1)
	}

	if ipRange == nil {
		os.Exit(1)
	}
	ips, err := parseIp(ipRange)
	if err != nil {
		os.Exit(1)
	}

	ports, err := parsePort(portRange)

	// fmt.Println("ip addr: ", ips)
	// fmt.Println("port number: ", ports)
	scan := scanfactory.NewScan(ips, ports)
	scan.Scan("connect")

	results := scan.GetResults()
	fmt.Println("\n----------------------------------------\n")
	fmt.Println("\nAll open port:\n")
	for _, result := range results {
		if result.Open {
			fmt.Printf("ip: %s port: %d open\n", result.Addr.Ip, result.Addr.Port)
		}
	}
}
