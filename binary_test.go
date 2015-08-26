package fiputil

import "testing"

func TestUint32ToByte4(t *testing.T) {
	i := uint32(127)
	b := UInt32ToByte4(i)
	t.Log(b)

	i1 := Byte4ToUint32(b)
	t.Log(i1)

	i = uint32(12345)
	b = UInt32ToByte4(i)
	t.Log(b)

	i1 = Byte4ToUint32(b)
	t.Log(i1)
}

func TestUintToByte(t *testing.T) {
	i := byte(2)
	b := UintToByte(uint32(i), uint8(1))
	t.Logf("%v", b)

	i1 := uint32(123456)
	b1 := UintToByte(i1, uint8(4))
	t.Logf("%v", b1)

	i2 := ByteToUint(b)
	t.Logf("%v", i2)
}
