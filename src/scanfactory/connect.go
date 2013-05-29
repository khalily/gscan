package scanfactory

import (
	"net"
	"strconv"
	"time"
	. "types"
)

func connect(addr Addr, c chan Result, ok chan bool) {
	remote := string(addr.Ip) + ":" + strconv.Itoa(int(addr.Port))
	// fmt.Println(remote)

	conn, err := net.DialTimeout("tcp", remote, 5*time.Second)
	// conn, err := net.Dial("tcp", remote)
	var result Result
	// fmt.Println("error:", err)
	if err == nil {
		conn.Close()
		result = Result{addr, true}
	} else {
		result = Result{addr, false}
		// fmt.Println("error: ", err)
	}
	c <- result
	ok <- true
}
