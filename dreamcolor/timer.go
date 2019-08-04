package dreamcolor

import "time"

type AutoTimerCommand struct {
	Start HourMinute
	End   HourMinute
}

type DelayCommand struct {
	Enable bool
	HourMinute
}

type TimeCommand struct {
	HourMinute
	Seconds   int
	DayOfWeek time.Weekday
}

func GetAutoTimer() []byte {
	return buildReadCommand(commandAutoTimer).toByteArray()
}

func SetAutoTimer(parameters AutoTimerCommand) []byte {
	return buildWriteCommand(commandAutoTimer).
		writeTime(parameters.Start).
		writeTime(parameters.End).
		toByteArray()
}

func GetDelay() []byte {
	return buildReadCommand(commandDelay).toByteArray()
}

func SetDelay(parameters DelayCommand) []byte {
	return buildWriteCommand(commandDelay).
		writeBoolean(parameters.Enable).
		writeTime(parameters.HourMinute).
		toByteArray()
}

func GetTime() []byte {
	return buildReadCommand(commandSync).toByteArray()
}

func SetTime(parameters TimeCommand) []byte {
	return buildWriteCommand(commandSync).
		writeTime(parameters.HourMinute).
		writeByte(parameters.Seconds).
		writeByte(int(parameters.DayOfWeek)).
		toByteArray()
}

func SyncTime() []byte {

	dateTime := time.Now()
	dayOfWeek := dateTime.Weekday()

	if dayOfWeek == time.Sunday {
		dayOfWeek = 7
	}

	return SetTime(TimeCommand{
		HourMinute{dateTime.Hour(), dateTime.Minute()},
		dateTime.Second(),
		dayOfWeek,
	})
}
