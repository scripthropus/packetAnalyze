package main

import (
	"fmt"
	"packetanalyze/packet"
	"syscall"
)

func htons(i uint16) uint16 {
	return (i << 8) | (i >> 8)
}

func main() {
	fd, err := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_RAW, int(htons(syscall.ETH_P_IP)))
	if err != nil {
		panic(err)
	}

	defer syscall.Close(fd)

	buf := make([]byte, 65535)
	for {
		n, _, err := syscall.Recvfrom(fd, buf, 0)
		if err != nil {
			panic(err)
		}

		ipv4 := packet.ParseIPv4Header(buf[14:n])
		fmt.Printf("%#v\n", ipv4)
	}

}
