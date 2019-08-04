package dreamcolor

const (
	SubmodeEnergy   = 0
	SubmodeSpectrum = 1
	SubmodeScroll   = 2
	SubmodeRhythm   = 3
	SubmodeMild     = 4
	SubmodeDynamic  = 5
)

type MusicCommand struct {
	Sensitivity int
}

type MusicRgbCommand struct {
	MusicCommand
	RgbColor
}

func BuildMusicCommand(submode int, sensitivity int) *Buffer {
	return BuildWriteCommand().
		WriteByte(CommandMode).
		WriteByte(ModeMusicIc).
		WriteByte(submode).
		WriteByte(sensitivity)
}

func SetDynamic(parameters MusicCommand) *Buffer {
	return BuildMusicCommand(SubmodeDynamic, parameters.Sensitivity)
}

func SetEnergy(parameters MusicCommand) *Buffer {
	return BuildMusicCommand(SubmodeEnergy, parameters.Sensitivity)
}

func SetMild(parameters MusicCommand) *Buffer {
	return BuildMusicCommand(SubmodeMild, parameters.Sensitivity)
}

func SetRhythm(parameters MusicCommand) *Buffer {
	return BuildMusicCommand(SubmodeRhythm, parameters.Sensitivity)
}

func SetScroll(parameters MusicRgbCommand) *Buffer {
	return BuildMusicCommand(SubmodeScroll, parameters.Sensitivity).WriteRgb(parameters.RgbColor)
}

func SetSpectrum(parameters MusicRgbCommand) *Buffer {
	return BuildMusicCommand(SubmodeSpectrum, parameters.Sensitivity).WriteRgb(parameters.RgbColor)
}
