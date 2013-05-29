package types

type Ip string
type Port uint16

type Addr struct {
	Ip   Ip
	Port Port
}

type Result struct {
	Addr Addr
	Open bool
}
