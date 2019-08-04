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

func BuildReadCommand() *Buffer {
	return InitializeBuffer(CommandLength).WriteByte(OpcodeRead)
}

func BuildWriteCommand() *Buffer {
	return InitializeBuffer(CommandLength).WriteByte(OpcodeWrite)
}

func BuildWriteBufferedCommand(parameters []byte) *Buffer {
	return InitializeBuffer(CommandLength).WriteByte(OpcodeWriteBuffered)
}
