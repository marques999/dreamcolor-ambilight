package dreamcolor

type HourMinute struct {
	Hours   int
	Minutes int
}

type RgbColor struct {
	Red   uint32
	Green uint32
	Blue  uint32
}

func (buffer *arrayBuffer) writeTime(time HourMinute) *arrayBuffer {
	return buffer.writeByte(time.Hours).writeByte(time.Minutes)
}

func (buffer *arrayBuffer) writeRgb(color RgbColor) *arrayBuffer {
	return buffer.writeByte(int(color.Red)).writeByte(int(color.Green)).writeByte(int(color.Blue))
}
