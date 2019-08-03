package main

import "time"

const COMMAND_TOGGLE = 0x01
const COMMAND_BRIGHTNESS = 0x04
const COMMAND_MODE = 0x05
const COMMAND_VERSION = 0x06
const COMMAND_CLOCK = 0x09
const COMMAND_TIMER = 0x0A

const MODE_MUSIC = 0x01
const MODE_STATIC = 0x02
const MODE_SCENE = 0x04
const MODE_PRESET = 0x0A

const MUSIC_ENERGETIC = 0x00
const MUSIC_SPECTRUM = 0x01
const MUSIC_RHYTHM = 0x03

func generateCommand(parameters []byte) []byte {
	return append(
		append(parameters, make([]byte, 19-len(parameters))...),
		calculateXor(parameters))
}

func calculateXor(parameters []byte) byte {

	var xorChecksum byte

	for _, parameter := range parameters {
		xorChecksum ^= parameter
	}

	return byte(xorChecksum & 0xFF)
}

func setEnabled(enabled bool) []byte {

	var enabledInt byte

	if enabled {
		enabledInt = 1
	}

	return []byte{COMMAND_TOGGLE, enabledInt}
}

func setBrightness(brightness byte) []byte {
	return []byte{COMMAND_BRIGHTNESS, brightness}
}

func setScene(sceneID byte) []byte {
	return []byte{COMMAND_MODE, MODE_SCENE, sceneID}
}

func setRhythm() []byte {
	return []byte{COMMAND_MODE, MODE_MUSIC, MUSIC_RHYTHM}
}

func setEnergetic() []byte {
	return []byte{COMMAND_MODE, MODE_MUSIC, MUSIC_ENERGETIC}
}

func setClock(dateTime time.Time) []byte {
	return []byte{COMMAND_CLOCK, uint8(dateTime.Hour()), uint8(dateTime.Minute()), uint8(dateTime.Second())}
}

func setColor(red uint32, green uint32, blue uint32) []byte {
	return []byte{COMMAND_MODE, MODE_STATIC, byte(red & 0xFF), byte(green & 0xFF), byte(blue & 0xFF)}
}

func setSpectrum(red byte, green byte, blue byte) []byte {
	return []byte{COMMAND_MODE, MODE_MUSIC, MUSIC_SPECTRUM, 0x00, red, green, blue}
}

func setPresetMode() []byte {
	return []byte{COMMAND_MODE, MODE_PRESET, 0x12}
}

// func generatePresetCommand(parameters []byte) []byte {
// 	return generateCommand(append([]byte{0xA1, 0x02}, parameters...))
// }

// func setPresetData(effect byte, zone byte, speed byte, colors [][3]byte) [][]byte {

// 	var packets [][]byte

// 	for _, packet := range generatePresetData(effect, zone, speed, colors) {
// 		packets = append(packets, generatePresetCommand(packet))
// 	}

// 	return packet
// }
