package dreamcolor

const (
	commandSwitch     = 0x01
	commandBrightness = 0x04
	commandMode       = 0x05
	commandVersion    = 0x06
	commandInitialize = 0x08
	commandSync       = 0x09
	commandAutoTimer  = 0x0A
	commandDelay      = 0x0B
	commandRedBlue    = 0x0D
	commandCurrent    = 0x0E
	commandLength     = 0x14
)

func buildReadCommand(address int) *arrayBuffer {
	return initializeBuffer(commandLength).writeByte(0xAA).writeByte(address)
}

func buildWriteCommand(address int) *arrayBuffer {
	return initializeBuffer(commandLength).writeByte(0x33).writeByte(address)
}

func buildWriteBufferedCommand(address int) *arrayBuffer {
	return initializeBuffer(commandLength).writeByte(0xA1).writeByte(address)
}
