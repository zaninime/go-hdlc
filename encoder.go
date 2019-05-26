package hdlc

import (
	"bytes"
	"encoding/binary"
	"io"

	"github.com/sigurn/crc16"
)

// An Encoder writes HDLC frames to an output stream.
type Encoder struct {
	w io.Writer
}

// NewEncoder returns a new encoder that writes to w.
func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{
		w: w,
	}
}

// WriteFrame writes the frame f on the output stream, encoding it's content.
func (e Encoder) WriteFrame(f *Frame) (int64, error) {
	var frameBuf bytes.Buffer
	var dataBuf bytes.Buffer

	if f.HasAddressCtrlPrefix {
		dataBuf.Write(addressCtrlSeq)
	}
	dataBuf.Write(f.Payload)
	dataBuf.Write(f.FCS)
	data := dataBuf.Bytes()

	frameBuf.WriteByte(flagSym)
	frameBuf.Write(escapeData(data))
	frameBuf.WriteByte(flagSym)

	return frameBuf.WriteTo(e.w)
}

func escapeData(p []byte) []byte {
	var out bytes.Buffer

	for _, b := range p {
		if (b) < 0x20 || ((b)&0x7f) == 0x7d || ((b)&0x7f) == 0x7e {
			out.WriteByte(escapeSym)
			out.WriteByte(b ^ 0x20)
		} else {
			out.WriteByte(b)
		}
	}

	return out.Bytes()
}

// Encapsulate takes a payload p and some configuration and creates a frame that
// can be written with an Encoder.
func Encapsulate(p []byte, hasAddressCtrlPrefix bool) *Frame {
	crc := crc16.Init(crcTable)

	if hasAddressCtrlPrefix {
		crc = crc16.Update(crc, addressCtrlSeq, crcTable)
	}

	crc = crc16.Update(crc, p, crcTable)
	crc = crc16.Complete(crc, crcTable)
	crc ^= 0xffff
	crcBytes := []byte{0, 0}
	binary.LittleEndian.PutUint16(crcBytes, crc)

	return &Frame{
		Payload:              p,
		FCS:                  crcBytes,
		HasAddressCtrlPrefix: hasAddressCtrlPrefix,
	}
}
