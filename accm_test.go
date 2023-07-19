package hdlc

import (
	"testing"
)

func Test_asyncControlCharacterMap_contains(t *testing.T) {
	tests := []struct {
		name string
		accm asyncControlCharacterMap
		char byte
		want bool
	}{
		{
			name: "first control character",
			accm: asyncControlCharacterMap(CharNUL),
			char: 0x00,
			want: true,
		},
		{
			name: "last control character",
			accm: asyncControlCharacterMap(CharUS),
			char: 0x1F,
			want: true,
		},
		{
			name: "non-control character",
			accm: asyncControlCharacterMap(CharALL),
			char: 0x20,
			want: false,
		},
		{
			name: "accm is uninitialized",
			char: '\n',
			want: false,
		},
		{
			name: "accm contains only the one character being checked",
			accm: asyncControlCharacterMap(CharLF),
			char: '\n',
			want: true,
		},
		{
			name: "accm contains all chars including the one being checked",
			accm: asyncControlCharacterMap(CharALL),
			char: '\t',
			want: true,
		},
		{
			name: "accm contains all chars except the one being checked",
			accm: asyncControlCharacterMap(CharALL &^ CharHT),
			char: '\t',
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.accm.contains(tt.char); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
