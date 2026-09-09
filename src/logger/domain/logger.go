package domain

type LogLevel int

const (
	Debug LogLevel = iota
	Info
	Warn
	Error
)

func (l LogLevel) String() string {
	switch l {
	case Debug:
		return "Debug"
	case Info:
		return "Info"
	case Warn:
		return "Warn"
	case Error:
		return "Error"
	default:
		return "Unknown"
	}
}

type ILog = func(level LogLevel, str ...string)

type IFormat = func(level LogLevel, str ...string) string

type ILogger interface {
	Error(err error, msg ...string)
	Warn(msg ...string)
	Info(msg ...string)
	Debug(msg ...string)
}
