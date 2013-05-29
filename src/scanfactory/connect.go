package scanfactory

import (
	"fmt"
	"net"
	"strconv"
	. "types"
)

func connect(addr Addr, c chan Result) {
	remote := string(addr.Ip) + ":" + strconv.Itoa(int(addr.Port))
	fmt.Println(remote)

	conn, err := net.Dial("tcp", remote)
	conn.Close()
	var result Result
	fmt.Println("error:", err)
	if err == nil {
		result = Result{addr, true}
	} else {
		result = Result{addr, false}
	}
	c <- result
}
