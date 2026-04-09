package main

import (
	"fmt"
	"packetanalyze/packet"
	"syscall"
)

func main() {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_ICMP)
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

		ipv4 := packet.ParseIPv4Header(buf[:n])
		fmt.Printf("%#v\n", ipv4)
	}

}
