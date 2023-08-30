package levels

type LogLevel int

type LogLevelType interface {
	~int
}

const LevelCnt = 8

const (
	None LogLevel = iota
	Emergency
	Crit
	Error
	Alert
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
	case Crit:
		return "Crit"
	case Error:
		return "Error"
	case Alert:
		return "Alert"
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
