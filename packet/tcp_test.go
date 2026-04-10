package packet

import (
	"encoding/hex"
	"reflect"
	"testing"
)

func TestParseTCPHeader(t *testing.T) {
	rawPacket, _ := hex.DecodeString("f3e1005087a9d6f315dfa8a9501800ffc21f0000")
	tests := []struct {
		name    string
		input   []byte
		want    *TCPHeader
		wantErr bool
	}{
		{
			name:  "tcp packt",
			input: rawPacket,
			want: &TCPHeader{
				SrcP:    0xf3e1,
				DistP:   0x0050,
				SeqNum:  0x87a9d6f3,
				AckNum:  0x15dfa8a9,
				DataOfs: 5,
				Rsrvd:   0x0000,
				CWR:     0,
				ECE:     0,
				URG:     0,
				ACK:     1,
				PSH:     1,
				RST:     0,
				SYN:     0,
				FIN:     0,
				WinSize: 0x00ff,
				ChSum:   0xc21f,
				UrgPtr:  0,
				Opt:     nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseTCPHeader(tt.input)
			if !reflect.DeepEqual(*got, *tt.want) {
				t.Errorf("got = %v, want = %v", got, *tt.want)
			}
		})
	}

}
