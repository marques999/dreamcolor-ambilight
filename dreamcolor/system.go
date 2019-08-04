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

func GetBrightness() []byte {
	return buildReadCommand(commandBrightness).toByteArray()
}

func GetEnabled() []byte {
	return buildReadCommand(commandSwitch).toByteArray()
}

func GetRedBlue() []byte {
	return buildReadCommand(commandRedBlue).toByteArray()
}

func GetVersion() []byte {
	return buildReadCommand(commandVersion).toByteArray()
}

func SetBrightness(parameters IntegerCommand) []byte {
	return buildWriteCommand(commandBrightness).
		writeByte(parameters.Value).
		toByteArray()
}

func SetCurrent(parameters IntegerCommand) []byte {
	return buildWriteCommand(commandCurrent).
		writeByte(parameters.Value).
		toByteArray()
}

func SetEnabled(parameters BooleanCommand) []byte {
	return buildWriteCommand(commandSwitch).
		writeBoolean(parameters.Value).
		toByteArray()
}

func SetInitialize() []byte {
	return buildWriteCommand(commandInitialize).toByteArray()
}

func SetRedBlue(parameters RedBlueCommand) []byte {
	return buildWriteCommand(commandRedBlue).
		writeByte(parameters.Red & 0xFF).
		writeByte(parameters.Blue & 0xFF).
		toByteArray()
}
