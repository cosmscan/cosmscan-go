package log

// Interface is a minimal interface for any logger
type Interface interface {
	Log(keyvals ...interface{})
}
