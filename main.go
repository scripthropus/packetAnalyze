package main

import (
	"fmt"
	"net"
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

		src := net.IP(buf[12:16])
		dst := net.IP(buf[16:20])
		fmt.Printf("src=%s dst=%s raw=%x\n", src, dst, buf[:n])
	}

}
