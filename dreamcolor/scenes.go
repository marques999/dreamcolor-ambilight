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

func SetScene(parameters IntegerCommand) *Buffer {
	return BuildWriteCommand(CommandMode).
		WriteByte(ModeScene).
		WriteByte(parameters.Value)
}
