package lang

import (
	"bytes"
	"encoding/binary"
	"unicode/utf16"
)

type String string

func (s String) ToUCS2Bytes() []byte {
	var b bytes.Buffer
	_ = binary.Write(&b, binary.BigEndian, s.ToCharArray())
	return b.Bytes()
}

func (s String) ToCharArray() []uint16 {
	return utf16.Encode([]rune(s))
}

func (s String) String() string {
	return string(s)
}

func (s String) GoString() string {
	return string(s)
}
