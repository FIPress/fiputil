package fiputil

import (
	"encoding/binary"
)

func UInt32ToByte4(i uint32) (b []byte) {
	b = make([]byte, 4)
	binary.BigEndian.PutUint32(b, i)
	return
}

func Byte4ToUint32(b []byte) (i uint32) {
	return binary.BigEndian.Uint32(b)
}

func UintToByte(i uint32, len uint8) (b []byte) {
	b = make([]byte, len)
	switch len {
	case 1:
		b[0] = byte(i)
	case 2:
		binary.BigEndian.PutUint16(b, uint16(i))
	case 4:
		binary.BigEndian.PutUint32(b, i)
	default:
		b = nil
	}
	return
}

func ByteToUint(b []byte) (i uint32) {
	l := len(b)
	switch l {
	case 1:
		return uint32(b[0])
	case 2:
		return uint32(binary.BigEndian.Uint16(b))
	case 4:
		return binary.BigEndian.Uint32(b)
	default:
		return uint32(0)
	}
}
