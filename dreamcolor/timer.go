package dreamcolor

import "time"

type TimeCommand struct {
	Hours     int
	Minutes   int
	Seconds   int
	DayOfWeek time.Weekday
}

type DelayCommand struct {
	Enable  bool
	Hours   int
	Minutes int
}

func SyncTime() *Buffer {

	dateTime := time.Now()
	dayOfWeek := dateTime.Weekday()

	if dayOfWeek == time.Sunday {
		dayOfWeek = 7
	}

	return SetTime(TimeCommand{
		dateTime.Hour(),
		dateTime.Minute(),
		dateTime.Second(),
		dayOfWeek,
	})
}

func SetDelay(parameters DelayCommand) *Buffer {

	return BuildWriteCommand().
		WriteByte(CommandDelay).
		WriteBoolean(parameters.Enable).
		WriteByte(parameters.Hours).
		WriteByte(parameters.Minutes)
}

func SetTime(parameters TimeCommand) *Buffer {

	return BuildWriteCommand().
		WriteByte(CommandSync).
		WriteByte(parameters.Hours).
		WriteByte(parameters.Minutes).
		WriteByte(parameters.Seconds).
		WriteByte(int(parameters.DayOfWeek))
}