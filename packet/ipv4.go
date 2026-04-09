package packet

import (
	"encoding/binary"
)

type IPv4Header struct {
	Version byte //上位4bit
	IHL     byte //下位4bit
	ToS     byte
	TLen    uint16
	ID      uint16
	Flag    byte   //上位3bit
	FO      uint16 //下位13bit
	TTL     byte
	Prot    byte
	HCS     uint16
	Src     uint32
	Dist    uint32
	Opt     []byte
}

func ParseIPv4Header(b []byte) *IPv4Header {
	ipv4 := IPv4Header{}
	ipv4.Version = (b[0] >> 4) & 0b1111
	ipv4.IHL = b[0] & 0b1111
	ipv4.ToS = b[1]
	ipv4.TLen = binary.BigEndian.Uint16(b[2:4])
	ipv4.ID = binary.BigEndian.Uint16(b[4:6])
	ipv4.Flag = (b[7] >> 5) & 0b111
	ipv4.FO = binary.BigEndian.Uint16(b[6:8]) & 0b1111111111111
	ipv4.TTL = b[8]
	ipv4.Prot = b[9]
	ipv4.HCS = binary.BigEndian.Uint16(b[10:12])
	ipv4.Src = binary.BigEndian.Uint32(b[12:16])
	ipv4.Dist = binary.BigEndian.Uint32(b[16:20])
	headerLen := int(ipv4.IHL) * 4
	if headerLen > 20 {
		ipv4.Opt = b[20:headerLen]
	}
	return &ipv4
}
