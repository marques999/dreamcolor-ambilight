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

func GetBrightness() *Buffer {
	return BuildReadCommand(CommandBrightness)
}

func GetEnabled() *Buffer {
	return BuildReadCommand(CommandSwitch)
}

func SetBrightness(parameters IntegerCommand) *Buffer {
	return BuildWriteCommand(CommandBrightness).WriteByte(parameters.Value)
}

func SetCurrent(parameters IntegerCommand) *Buffer {
	return BuildWriteCommand(CommandCurrent).WriteByte(parameters.Value)
}

func SetEnabled(parameters BooleanCommand) *Buffer {
	return BuildWriteCommand(CommandSwitch).WriteBoolean(parameters.Value)
}

func SetInitialize() *Buffer {
	return BuildWriteCommand(CommandInitialize)
}

func SetRedBlue(parameters RedBlueCommand) *Buffer {
	return BuildWriteCommand(CommandRedBlue).
		WriteByte(parameters.Red & 0xFF).
		WriteByte(parameters.Blue & 0xFF)
}
