package packet

import (
	"encoding/binary"
)

type TCPHeader struct {
	SrcP    uint16
	DistP   uint16
	SeqNum  uint32
	AckNum  uint32
	DataOfs byte //上位4bit
	Rsrvd   byte //下位4bit
	//Flags
	CWR     byte //1bit目
	ECE     byte //2bit目
	URG     byte //3bit目
	ACK     byte //4bit目
	PSH     byte //5bit目
	RST     byte //6bit目
	SYN     byte //7bit目
	FIN     byte //8bit目
	WinSize uint16
	ChSum   uint16
	UrgPtr  uint16
	Opt     []byte
}

func ParseTCPHeader(b []byte) *TCPHeader {
	tcp := TCPHeader{}
	tcp.SrcP = binary.BigEndian.Uint16(b[0:2])
	tcp.DistP = binary.BigEndian.Uint16(b[2:4])
	tcp.SeqNum = binary.BigEndian.Uint32(b[4:8])
	tcp.AckNum = binary.BigEndian.Uint32(b[8:12])
	tcp.DataOfs = (b[12] >> 4) & 0b1111
	tcp.Rsrvd = b[12] & 0b1111
	//flags b[13]

	tcp.CWR = (b[13] >> 7) & 1
	tcp.ECE = (b[13] >> 6) & 1
	tcp.URG = (b[13] >> 5) & 1
	tcp.ACK = (b[13] >> 4) & 1
	tcp.PSH = (b[13] >> 3) & 1
	tcp.RST = (b[13] >> 2) & 1
	tcp.SYN = (b[13] >> 1) & 1
	tcp.FIN = (b[13] >> 0) & 1
	tcp.WinSize = binary.BigEndian.Uint16(b[14:16])
	tcp.ChSum = binary.BigEndian.Uint16(b[16:18])
	optSize := 4*tcp.DataOfs - 20
	if optSize > 20 {
		tcp.Opt = b[20:optSize]
	}

	return &tcp
}
