package dreamcolor

type arrayBuffer struct {
	offset   int
	capacity int
	data     []byte
}

func calculateXor(parameters []byte) byte {

	var xorChecksum byte

	for _, parameter := range parameters {
		xorChecksum ^= parameter
	}

	return xorChecksum
}

func initializeBuffer(capacity int) *arrayBuffer {
	return &arrayBuffer{
		offset:   0,
		capacity: capacity,
		data:     make([]byte, capacity),
	}
}

func (writer *arrayBuffer) writeByte(value int) *arrayBuffer {

	writer.data[writer.offset] = byte(value & 0xFF)
	writer.offset++

	return writer
}

func (writer *arrayBuffer) writeBoolean(value bool) *arrayBuffer {

	if value {
		return writer.writeByte(1)
	}

	return writer.writeByte(0)
}

func (writer *arrayBuffer) toByteArray() []byte {

	writer.data[writer.capacity-1] = calculateXor(writer.data)

	return writer.data
}
