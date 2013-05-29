package scanfactory

import (
	"fmt"
	. "types"
)

type Addrs struct {
	ipList   []Ip
	portList []Port
}

type Scan struct {
	Addrs
	chans   []chan Result
	results []Result
}

func NewScan(ipList []Ip, portList []Port) (scan *Scan) {
	sumAddr := len(ipList) * len(portList)
	chans := make([]chan Result, sumAddr)
	return &Scan{Addrs: Addrs{ipList, portList}, chans: chans}
}

func (s *Scan) Scan(way string) {
	switch way {
	case "connect":
		for i, ip := range s.ipList {
			for j, port := range s.portList {
				addr := Addr{ip, port}
				c := make(chan Result, 0)
				s.chans[i+j] = c

				go connect(addr, c)
			}
		}
	case "syn":
		fmt.Println("sorry, syn not support!")
		return
	}

	s.waitResults()
}

func (s *Scan) waitResults() {
	for _, c := range s.chans {
		result := <-c
		s.results = append(s.results, result)
		fmt.Println("result:", result)
	}
}

func (s *Scan) GetResults() []Result {
	return s.results
}

func (s *Scan) GetAddr() {
	for _, ip := range s.ipList {
		fmt.Println(ip)
	}

	for port := range s.portList {
		fmt.Println(port)
	}
}
