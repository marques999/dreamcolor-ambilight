package dreamcolor

const (
	ModeVideo      = 0x00
	ModeMusicIc    = 0x01
	ModeColor      = 0x02
	ModeMusic      = 0x03
	ModeScene      = 0x04
	ModeMicrophone = 0x05
	ModeDiySingle  = 0x06
	ModeStatic7001 = 0x08
	ModeScene7001  = 0x09
	ModeDiy        = 0x0A
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

func GetMode() *Buffer {
	return BuildReadCommand().WriteByte(CommandMode)
}

func SetColor(parameters RgbColor) *Buffer {
	return BuildWriteCommand().WriteByte(CommandMode).WriteByte(ModeColor)
}

func SetColorAlternate(parameters ColorCommand) *Buffer {
	return SetColor(parameters.ColorA).
		WriteBoolean(parameters.Toggle).
		WriteRgb(parameters.ColorB)
}

func SetMicrophone(parameters MicrophoneCommand) *Buffer {
	return BuildWriteCommand().
		WriteByte(CommandMode).
		WriteByte(ModeMicrophone).
		WriteBoolean(parameters.Toggle).
		WriteRgb(parameters.RgbColor)
}
