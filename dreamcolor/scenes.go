package dreamcolor

type Scene int

var Scenes = &struct {
	Morning  Scene
	Sunset   Scene
	Movie    Scene
	Date     Scene
	Romantic Scene
	Blinking Scene
	Candle   Scene
	Snow     Scene
}{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07}

type SceneCommand struct {
	Id Scene
}

func SetScene(parameters SceneCommand) []byte {
	return buildWriteCommand(commandMode).
		writeByte(modeScene).
		writeByte(int(parameters.Id)).
		toByteArray()
}
