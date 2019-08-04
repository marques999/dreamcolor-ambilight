package dreamcolor

const (
	submodeEnergy   = 0
	submodeSpectrum = 1
	submodeScroll   = 2
	submodeRhythm   = 3
	submodeMild     = 4
	submodeDynamic  = 5
)

type MusicCommand struct {
	Sensitivity int
}

type MusicRgbCommand struct {
	MusicCommand
	RgbColor
}

func writeMusic(submode int, sensitivity int) *arrayBuffer {
	return buildWriteCommand(commandMode).
		writeByte(modeMusicIc).
		writeByte(submode).
		writeByte(sensitivity)
}

func SetDynamic(parameters MusicCommand) []byte {
	return writeMusic(submodeDynamic, parameters.Sensitivity).toByteArray()
}

func SetEnergy(parameters MusicCommand) []byte {
	return writeMusic(submodeEnergy, parameters.Sensitivity).toByteArray()
}

func SetMild(parameters MusicCommand) []byte {
	return writeMusic(submodeMild, parameters.Sensitivity).toByteArray()
}

func SetRhythm(parameters MusicCommand) []byte {
	return writeMusic(submodeRhythm, parameters.Sensitivity).toByteArray()
}

func SetScroll(parameters MusicRgbCommand) []byte {
	return writeMusic(submodeScroll, parameters.Sensitivity).
		writeRgb(parameters.RgbColor).
		toByteArray()
}

func SetSpectrum(parameters MusicRgbCommand) []byte {
	return writeMusic(submodeSpectrum, parameters.Sensitivity).
		writeRgb(parameters.RgbColor).
		toByteArray()
}
