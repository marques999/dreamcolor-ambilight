package dreamcolor

const (
	CommandSwitch     = 0x01
	CommandBrightness = 0x04
	CommandMode       = 0x05
	CommandVersion    = 0x06
	CommandInitialize = 0x08
	CommandSync       = 0x09
	CommandTimer      = 0x0A
	CommandDelay      = 0x0B
	CommandRedBlue    = 0x0D
	CommandCurrent    = 0x0E
	CommandLength     = 0x14
)

const (
	OpcodeRead          = 0xAA
	OpcodeWrite         = 0x33
	OpcodeWriteBuffered = 0xA1
)

func BuildReadCommand(address int) *Buffer {
	return InitializeBuffer(CommandLength).WriteByte(OpcodeRead).WriteByte(address)
}

func BuildWriteCommand(address int) *Buffer {
	return InitializeBuffer(CommandLength).WriteByte(OpcodeWrite).WriteByte(address)
}

func BuildWriteBufferedCommand(address int) *Buffer {
	return InitializeBuffer(CommandLength).WriteByte(OpcodeWriteBuffered).WriteByte(address)
}
