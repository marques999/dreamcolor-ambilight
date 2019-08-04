package dreamcolor

type BooleanCommand struct {
	Value bool
}

type IntegerCommand struct {
	Value int
}

type RedBlueCommand struct {
	Red  int
	Blue int
}

func GetEnabled() *Buffer {
	return BuildReadCommand().WriteByte(CommandSwitch)
}

func SetBrightness(parameters IntegerCommand) *Buffer {
	return BuildWriteCommand().
		WriteByte(CommandBrightness).
		WriteByte(parameters.Value)
}

func SetCurrent(parameters IntegerCommand) *Buffer {
	return BuildWriteCommand().
		WriteByte(CommandCurrent).
		WriteByte(parameters.Value)
}

func SetEnabled(parameters BooleanCommand) *Buffer {
	return BuildWriteCommand().
		WriteByte(CommandSwitch).
		WriteBoolean(parameters.Value)
}

func SetInitialize() *Buffer {
	return BuildWriteCommand().WriteByte(CommandInitialize)
}

func SetRedBlue(parameters RedBlueCommand) *Buffer {
	return BuildWriteCommand().
		WriteByte(CommandRedBlue).
		WriteByte(parameters.Red & 0xFF).
		WriteByte(parameters.Blue & 0xFF)
}
