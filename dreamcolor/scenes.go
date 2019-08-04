package dreamcolor

const (
	SceneMorning  = 0x00
	SceneSunset   = 0x01
	SceneMovie    = 0x02
	SceneDate     = 0x03
	SceneRomantic = 0x04
	SceneBlinking = 0x05
	SceneCandle   = 0x06
	SceneSnow     = 0x07
)

type SceneCommand struct {
	ID int
}

func SetScene(parameters SceneCommand) *Buffer {
	return BuildWriteCommand().
		WriteByte(CommandMode).
		WriteByte(ModeScene).
		WriteByte(parameters.ID)
}
