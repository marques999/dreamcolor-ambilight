package dreamcolor

const (
	modeVideo      = 0x00
	modeMusicIc    = 0x01
	modeColor      = 0x02
	modeMusic      = 0x03
	modeScene      = 0x04
	modeMicrophone = 0x05
	modeDiySingle  = 0x06
	modeStatic7001 = 0x08
	modeScene7001  = 0x09
	modeDiy        = 0x0A
)

type ColorCommand struct {
	ColorA RgbColor
	ColorB RgbColor
	Toggle bool
}

type MicrophoneCommand struct {
	RgbColor
	Toggle bool
}

func GetMode() []byte {
	return buildReadCommand(commandMode).toByteArray()
}

func SetColor(parameters RgbColor) []byte {
	return buildWriteCommand(commandMode).
		writeByte(modeColor).
		writeRgb(parameters).
		toByteArray()
}

func SetColorAlternate(parameters ColorCommand) []byte {
	return buildWriteCommand(commandMode).
		writeByte(modeColor).
		writeRgb(parameters.ColorA).
		writeBoolean(parameters.Toggle).
		writeRgb(parameters.ColorB).
		toByteArray()
}

func SetMicrophone(parameters MicrophoneCommand) []byte {
	return buildWriteCommand(commandMode).
		writeByte(modeMicrophone).
		writeBoolean(parameters.Toggle).
		writeRgb(parameters.RgbColor).
		toByteArray()
}
