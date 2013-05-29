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
	ch      chan Result
	ok      chan bool
	results []Result
	sumAddr int
}

func NewScan(ipList []Ip, portList []Port) (scan *Scan) {
	sumAddr := len(ipList) * len(portList)
	ch := make(chan Result, sumAddr)
	ok := make(chan bool, sumAddr)
	return &Scan{Addrs: Addrs{ipList, portList}, ch: ch, ok: ok, sumAddr: sumAddr}
}

func (s *Scan) Scan(way string) {
	switch way {
	case "connect":
		for _, ip := range s.ipList {
			for _, port := range s.portList {
				addr := Addr{ip, port}

				go connect(addr, s.ch, s.ok)
			}
		}
	case "syn":
		fmt.Println("sorry, syn not support!")
		return
	}

	s.waitResults()
}

func (s *Scan) waitResults() {
	var result Result
	for {
		select {
		case result = <-s.ch:
			s.results = append(s.results, result)
			if result.Open {
				fmt.Printf("ip: %s port: %d open\n", result.Addr.Ip, result.Addr.Port)
			} else {
				// fmt.Printf("ip: %s port: %d not open\n", result.Addr.Ip, result.Addr.Port)
			}
			// for result = range s.ch {
			// 	s.results = append(s.results, result)
			// 	if result.Open {
			// 		fmt.Printf("ip: %s port: %d open\n", result.Addr.Ip, result.Addr.Port)
			// 	} else {
			// 		fmt.Printf("ip: %s port: %d not open\n", result.Addr.Ip, result.Addr.Port)
			// 	}
			// }
		case <-s.ok:
			s.sumAddr--
			// fmt.Println("sumAddr:", s.sumAddr)
			// for _ = range s.ok {
			// 	s.sumAddr--

			// 	fmt.Println("sumAddr:", s.sumAddr)
			// }
			if s.sumAddr == 0 {
				fmt.Println("over")
				goto LOOP
			}
		}
	}
LOOP:
}

func (s *Scan) GetResults() []Result {
	return s.results
}
