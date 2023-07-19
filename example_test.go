package hdlc_test

import (
	"fmt"
	"net"

	"github.com/zaninime/go-hdlc"
)

func Example() {
	conn1, conn2 := net.Pipe()

	// Decode
	frameCh := make(chan *hdlc.Frame)
	go func() {
		decoder := hdlc.NewDecoder(conn2)
		frame, err := decoder.ReadFrame()
		if err != nil {
			panic(err)
		}
		frameCh <- frame
	}()

	// Encode
	{
		encoder := hdlc.NewEncoder(conn1)
		frame := hdlc.Encapsulate([]byte("hello world"), false)
		if _, err := encoder.WriteFrame(frame); err != nil {
			panic(err)
		}
	}

	frame := <-frameCh

	fmt.Printf("%s", frame.Payload)
	// Output: hello world
}

// Register the Line Feed and Horizontal Tab characters in the
// Async-Control-Character-Map.
func Example_accm() {
	conn1, conn2 := net.Pipe()
	accm := hdlc.CharHT | hdlc.CharLF

	// Decode
	frameCh := make(chan *hdlc.Frame)
	go func() {
		decoder := hdlc.NewDecoder(conn2).SetACCM(accm)
		frame, err := decoder.ReadFrame()
		if err != nil {
			panic(err)
		}
		frameCh <- frame
	}()

	// Encode
	{
		encoder := hdlc.NewEncoder(conn1).SetACCM(accm)
		frame := hdlc.Encapsulate([]byte("hello\n\tworld"), false)
		if _, err := encoder.WriteFrame(frame); err != nil {
			panic(err)
		}
	}

	frame := <-frameCh

	fmt.Printf("%s", frame.Payload)
	// Output: hello
	//	world
}
