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

func BuildReadCommand(address int) *Buffer {
	return InitializeBuffer(CommandLength).WriteByte(0xAA).WriteByte(address)
}

func BuildWriteCommand(address int) *Buffer {
	return InitializeBuffer(CommandLength).WriteByte(0x33).WriteByte(address)
}

func BuildWriteBufferedCommand(address int) *Buffer {
	return InitializeBuffer(CommandLength).WriteByte(0xA1).WriteByte(address)
}
