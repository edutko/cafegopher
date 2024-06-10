package java

import (
	"encoding/binary"
	"io"
	"unicode/utf16"
)

type binaryReader struct {
	r io.Reader
}

func (r *binaryReader) readBytes(count int64) ([]byte, error) {
	b := make([]byte, count)
	_, err := io.ReadFull(r.r, b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (r *binaryReader) readByte() (byte, error) {
	b, err := r.readBytes(1)
	if err != nil {
		return 0, err
	}
	return b[0], nil
}

func (r *binaryReader) readBoolean() (bool, error) {
	b, err := r.readBytes(1)
	if err != nil {
		return false, err
	}
	return b[0] != 0, nil
}

func (r *binaryReader) readChar() (rune, error) {
	u, err := r.readUint16()
	if err != nil {
		return 0, err
	}
	return utf16.Decode([]uint16{u})[0], nil
}

func (r *binaryReader) readFloat32() (float32, error) {
	var f float32
	err := binary.Read(r.r, binary.BigEndian, &f)
	if err != nil {
		return 0, err
	}
	return f, nil
}

func (r *binaryReader) readFloat64() (float64, error) {
	var f float64
	err := binary.Read(r.r, binary.BigEndian, &f)
	if err != nil {
		return 0, err
	}
	return f, nil
}

func (r *binaryReader) readInt8() (int8, error) {
	b, err := r.readBytes(1)
	if err != nil {
		return 0, err
	}
	return int8(b[0]), nil
}

func (r *binaryReader) readInt16() (int16, error) {
	b, err := r.readBytes(2)
	if err != nil {
		return 0, err
	}
	return int16(binary.BigEndian.Uint16(b)), nil
}

func (r *binaryReader) readInt32() (int32, error) {
	b, err := r.readBytes(4)
	if err != nil {
		return 0, err
	}
	return int32(binary.BigEndian.Uint32(b)), nil
}

func (r *binaryReader) readInt64() (int64, error) {
	b, err := r.readBytes(8)
	if err != nil {
		return 0, err
	}
	return int64(binary.BigEndian.Uint64(b)), nil
}

func (r *binaryReader) readTypeCode() (TypeCode, error) {
	t, err := r.readByte()
	if err != nil {
		return 0, err
	}
	return TypeCode(t), nil
}

func (r *binaryReader) readUint8() (uint8, error) {
	return r.readByte()
}

func (r *binaryReader) readUint16() (uint16, error) {
	b, err := r.readBytes(2)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint16(b), nil
}
