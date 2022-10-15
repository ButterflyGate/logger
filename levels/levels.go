package levels

type LogLevel int

const (
	None LogLevel = iota
	Emergency
	Alert
	Critical
	Error
	Warning
	Notice
	Informational
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
	case Critical:
		return "Crit"
	case Error:
		return "Error"
	case Warning:
		return "Warn"
	case Notice:
		return "Notice"
	case Informational:
		return "Info"
	case Debug:
		return "Debug"
	case Trace:
		return "Trace"
	}
	return "Unknown"
}
