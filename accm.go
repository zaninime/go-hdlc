package hdlc

// ControlChar is an ASCII control character, which is a character less than
// 0x20.
//
// https://en.wikipedia.org/wiki/ASCII#Control_code_chart
type ControlChar uint32

const (
	CharNUL ControlChar = 1 << iota // Null
	CharSOH                         // Start of Head
	CharSTX                         // Start of Text
	CharETX                         // End of Text
	CharEOT                         // End of Transmission
	CharENQ                         // Enquiry
	CharACK                         // Acknowledgement
	CharBEL                         // Bell
	CharBS                          // Backspace
	CharHT                          // Horizontal Tab
	CharLF                          // Line Feed
	CharVT                          // vertical Tab
	CharFF                          // Form Feed
	CharCR                          // Carriage Return
	CharSO                          // Shift Out
	CharSI                          // Shift In
	CharDLE                         // Data Link Escape
	CharDC1                         // Device Control 1 (XON)
	CharDC2                         // Device Control 2
	CharDC3                         // Device Control 3 (XOFF)
	CharDC4                         // Device Control 4
	CharNAK                         // Negative Acknowledgement
	CharSYN                         // Synchronous Idle
	CharETB                         // End of Transmission Block
	CharCAN                         // Cancel
	CharEM                          // End of Medium
	CharSUB                         // Substitute
	CharESC                         // Escape
	CharFS                          // File Separator
	CharGS                          // Group Separator
	CharRS                          // Record Separator
	CharUS                          // Unit Separator

	// All control characters
	CharALL = CharNUL | CharSOH | CharSTX | CharETX | CharEOT | CharENQ |
		CharACK | CharBEL | CharBS | CharHT | CharLF | CharVT | CharFF | CharCR |
		CharSO | CharSI | CharDLE | CharDC1 | CharDC2 | CharDC3 | CharDC4 |
		CharNAK | CharSYN | CharETB | CharCAN | CharEM | CharSUB | CharESC |
		CharFS | CharGS | CharRS | CharUS
)

type asyncControlCharacterMap ControlChar

// contains returns true if accm contains the given control character and
// returns false otherwise.
func (accm *asyncControlCharacterMap) contains(char byte) bool {
	if char >= 0x20 {
		return false
	}
	return (*accm)&(1<<char) != 0
}
