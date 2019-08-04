package dreamcolor

import "encoding/hex"

type Buffer struct {
	data     []byte
	Capacity int
	Offset   int
}

func calculateXor(parameters []byte) byte {

	var xorChecksum byte

	for _, parameter := range parameters {
		xorChecksum ^= parameter
	}

	return byte(xorChecksum & 0xFF)
}

func InitializeBuffer(capacity int) *Buffer {
	return &Buffer{make([]byte, capacity), capacity, 0}
}

func (writer *Buffer) WriteByte(value int) *Buffer {

	writer.data[writer.Offset] = byte(value & 0xFF)
	writer.Offset++

	return writer
}

func (writer *Buffer) WriteBoolean(value bool) *Buffer {

	if value {
		return writer.WriteByte(0x01)
	}

	return writer.WriteByte(0x00)
}

func (writer *Buffer) Bytes() []byte {
	writer.data[writer.Capacity-1] = calculateXor(writer.data)
	return writer.data
}

func (writer *Buffer) String() string {
	return hex.EncodeToString(writer.data)
}
