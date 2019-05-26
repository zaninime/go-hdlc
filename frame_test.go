package hdlc

import (
	"bytes"
	"reflect"
	"testing"
)

func TestRound(t *testing.T) {
	payload := []byte("1234")
	t.Run("full round encapsulate/valid", func(t *testing.T) {
		frame := Encapsulate(payload, false)

		if !frame.Valid() {
			t.Error("created frame isn't valid")
		}
	})
	t.Run("full round encode/decode", func(t *testing.T) {
		var buf bytes.Buffer
		encoder := NewEncoder(&buf)
		_, err := encoder.WriteFrame(Encapsulate(payload, true))

		if err != nil {
			t.Error(err)
		}

		decoder := NewDecoder(&buf)
		frame, err := decoder.ReadFrame()

		if err != nil {
			t.Error(err)
		}

		if !frame.HasAddressCtrlPrefix {
			t.Error("frame was not recognized as having the addressCtrlPrefix")
		}

		if !reflect.DeepEqual(payload, frame.Payload) {
			t.Error("final payload doesn't match initial payload")
		}

		if !frame.Valid() {
			t.Error("final frame is not valid")
		}
	})
}

func TestFrame_Valid(t *testing.T) {
	tests := []struct {
		name string
		f    Frame
		want bool
	}{
		{
			name: "test vector payload",
			f: Frame{
				Payload:              []byte("123456789"),
				FCS:                  []byte{0x6e, 0x90}, // 0x6f91 ^ 0xffff, little endian
				HasAddressCtrlPrefix: false,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.Valid(); got != tt.want {
				t.Errorf("Frame.Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}
