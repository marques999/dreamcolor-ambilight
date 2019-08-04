package dreamcolor

type RgbColor struct {
	Red   uint32
	Green uint32
	Blue  uint32
}

func (writer *Buffer) WriteRgb(color RgbColor) *Buffer {
	return writer.
		WriteByte(int(color.Red)).
		WriteByte(int(color.Green)).
		WriteByte(int(color.Blue))
}
