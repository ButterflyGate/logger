package levels

type LogLevel int

const LevelCnt = 8

const (
	None LogLevel = iota
	Emergency
	Alert
	Crit
	Error
	Warn
	Notice
	Info
	Debug
	Trace LogLevel = 99
)

func (ll LogLevel) String() string {
	switch ll {
	case None:
		return "None"
	case Emergency:
		return "Emergency"
	case Alert:
		return "Alert"
	case Crit:
		return "Crit"
	case Error:
		return "Error"
	case Warn:
		return "Warn"
	case Notice:
		return "Notice"
	case Info:
		return "Info"
	case Debug:
		return "Debug"
	case Trace:
		return "Trace"
	}
	return "Unknown"
}
