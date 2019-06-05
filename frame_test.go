package hdlc

import (
	"bytes"
	"math/rand"
	"reflect"
	"strings"
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

func BenchmarkEncodeDecode(b *testing.B) {
	payloads := []struct {
		Description string
		PayloadSize int
	}{
		{
			Description: "Small",
			PayloadSize: 20,
		},
		{
			Description: "Medium",
			PayloadSize: 700,
		},
		{
			Description: "Big",
			PayloadSize: 1500,
		},
	}
	hasAddressCtrlPrefixes := []bool{true, false}

	for _, payloadDescr := range payloads {
		payload := make([]byte, payloadDescr.PayloadSize)

		for _, hasAddressCtrlPrefix := range hasAddressCtrlPrefixes {
			var name strings.Builder
			name.WriteString(payloadDescr.Description)
			name.WriteString("Payload")

			if hasAddressCtrlPrefix {
				name.WriteString("WithAddressCtrlPrefix")
			} else {
				name.WriteString("WithoutAddressCtrlPrefix")
			}

			var buf bytes.Buffer
			encoder := NewEncoder(&buf)
			decoder := NewDecoder(&buf)
			rand.Read(payload)
			b.Run(name.String(), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					b.SetBytes(int64(payloadDescr.PayloadSize))
					encoder.WriteFrame(Encapsulate(payload, true))
					decoder.ReadFrame()
				}
			})
		}
	}
}
