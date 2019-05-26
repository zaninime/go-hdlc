package hdlc

import (
	"reflect"
	"testing"
)

func TestEncapsulate(t *testing.T) {
	type args struct {
		p                    []byte
		hasAddressCtrlPrefix bool
	}
	tests := []struct {
		name string
		args args
		want *Frame
	}{
		{
			name: "real packet 1",
			args: args{
				p:                    []byte{0x08, 0x91},
				hasAddressCtrlPrefix: false,
			},
			want: &Frame{
				Payload:              []byte{0x08, 0x91},
				FCS:                  []byte{0x87, 0x44},
				HasAddressCtrlPrefix: false,
			},
		},
		{
			name: "real packet 2",
			args: args{
				p:                    []byte{0x08, 0xb1},
				hasAddressCtrlPrefix: false,
			},
			want: &Frame{
				Payload:              []byte{0x08, 0xb1},
				FCS:                  []byte{0x85, 0x65},
				HasAddressCtrlPrefix: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encapsulate(tt.args.p, tt.args.hasAddressCtrlPrefix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encapsulate() = %v, want %v", got, tt.want)
			}
		})
	}
}
