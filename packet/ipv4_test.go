package packet

import (
	"encoding/hex"
	"reflect"
	"testing"
)

func TestParseIPv4Header(t *testing.T) {
	rawPacket, _ := hex.DecodeString("45000054000000007301738a08080808c0a8036700002d00581c00157ff1d6690000000058a00d0000000000101112131415161718191a1b1c1d1e1f202122232425262728292a2b2c2d2e2f3031323334353637")

	tests := []struct {
		name    string
		input   []byte
		want    *IPv4Header
		wantErr bool
	}{
		{
			name:  "packet",
			input: rawPacket,
			want: &IPv4Header{
				Version: 4,
				IHL:     5,
				ToS:     0,
				TLen:    0x54,
				ID:      0,
				Flag:    0,
				FO:      0,
				TTL:     0x73,
				Prot:    0x01,
				HCS:     0x738a,
				Src:     0x08080808,
				Dist:    0xc0a80367,
				Opt:     nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseIPv4Header(tt.input)
			if !reflect.DeepEqual(*got, *tt.want) {
				t.Errorf("got = %v, want = %v", got, *tt.want)
			}
		})
	}

}
