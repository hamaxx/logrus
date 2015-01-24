package logrus

import (
	"io"
	"os"
	"sync"
)

type Logger struct {
	// The logs are `io.Copy`'d to this in a mutex. It's common to set this to a
	// file, or leave it default which is `os.Stdout`. You can also set this to
	// something more adventorous, such as logging to Kafka.
	Out io.Writer
	// Hooks for the logger instance. These allow firing events based on logging
	// levels and log entries. For example, to send errors to an error tracking
	// service, log to StatsD or dump the core on fatal errors.
	Hooks levelHooks
	// All log entries pass through the formatter before logged to Out. The
	// included formatters are `TextFormatter` and `JSONFormatter` for which
	// TextFormatter is the default. In development (when a TTY is attached) it
	// logs with colors, but to a file it wouldn't. You can easily implement your
	// own that implements the `Formatter` interface, see the `README` or included
	// formatters for examples.
	Formatter Formatter
	// The logging level the logger should log at. This is typically (and defaults
	// to) `logrus.Info`, which allows Info(), Warn(), Error() and Fatal() to be
	// logged. `logrus.Debug` is useful in
	Level Level
	// Used to sync writing to the log.
	mu sync.Mutex
}

// Creates a new logger. Configuration should be set by changing `Formatter`,
// `Out` and `Hooks` directly on the default logger instance. You can also just
// instantiate your own:
//
//    var log = &Logger{
//      Out: os.Stderr,
//      Formatter: new(JSONFormatter),
//      Hooks: make(levelHooks),
//      Level: logrus.DebugLevel,
//    }
//
// It's recommended to make this a global instance called `log`.
func New() *Logger {
	return &Logger{
		Out:       os.Stdout,
		Formatter: new(TextFormatter),
		Hooks:     make(levelHooks),
		Level:     InfoLevel,
	}
}

// Adds a field to the log entry, note that you it doesn't log until you call
// Debug, Print, Info, Warn, Fatal or Panic. It only creates a log entry.
// Ff you want multiple fields, use `WithFields`.
func (logger *Logger) WithField(key string, value interface{}) *Entry {
	return NewEntry(logger).WithField(key, value)
}

// Adds a struct of fields to the log entry. All it does is call `WithField` for
// each `Field`.
func (logger *Logger) WithFields(fields Fields) *Entry {
	return NewEntry(logger).WithFields(fields)
}

func (logger *Logger) Debugf(format string, args ...interface{}) {
	e := NewEntry(logger)
	e.Debugf(format, args...)
	entryPool.Put(e)
}

func (logger *Logger) Infof(format string, args ...interface{}) {
	e := NewEntry(logger)
	e.Infof(format, args...)
	entryPool.Put(e)
}

func (logger *Logger) Printf(format string, args ...interface{}) {
	e := NewEntry(logger)
	e.Printf(format, args...)
	entryPool.Put(e)
}

func (logger *Logger) Warnf(format string, args ...interface{}) {
	e := NewEntry(logger)
	e.Warnf(format, args...)
	entryPool.Put(e)
}

func (logger *Logger) Warningf(format string, args ...interface{}) {
	e := NewEntry(logger)
	e.Warnf(format, args...)
	entryPool.Put(e)
}

func (logger *Logger) Errorf(format string, args ...interface{}) {
	e := NewEntry(logger)
	e.Errorf(format, args...)
	entryPool.Put(e)
}

func (logger *Logger) Fatalf(format string, args ...interface{}) {
	e := NewEntry(logger)
	e.Fatalf(format, args...)
	entryPool.Put(e)
}

func (logger *Logger) Panicf(format string, args ...interface{}) {
	e := NewEntry(logger)
	e.Panicf(format, args...)
	entryPool.Put(e)
}

func (logger *Logger) Debug(args ...interface{}) {
	e := NewEntry(logger)
	e.Debug(args...)
	entryPool.Put(e)
}

func (logger *Logger) Info(args ...interface{}) {
	e := NewEntry(logger)
	e.Info(args...)
	entryPool.Put(e)
}

func (logger *Logger) Print(args ...interface{}) {
	e := NewEntry(logger)
	e.Info(args...)
	entryPool.Put(e)
}

func (logger *Logger) Warn(args ...interface{}) {
	e := NewEntry(logger)
	e.Warn(args...)
	entryPool.Put(e)
}

func (logger *Logger) Warning(args ...interface{}) {
	e := NewEntry(logger)
	e.Warn(args...)
	entryPool.Put(e)
}

func (logger *Logger) Error(args ...interface{}) {
	e := NewEntry(logger)
	e.Error(args...)
	entryPool.Put(e)
}

func (logger *Logger) Fatal(args ...interface{}) {
	e := NewEntry(logger)
	e.Fatal(args...)
	entryPool.Put(e)
}

func (logger *Logger) Panic(args ...interface{}) {
	e := NewEntry(logger)
	e.Panic(args...)
	entryPool.Put(e)
}

func (logger *Logger) Debugln(args ...interface{}) {
	e := NewEntry(logger)
	e.Debugln(args...)
	entryPool.Put(e)
}

func (logger *Logger) Infoln(args ...interface{}) {
	e := NewEntry(logger)
	e.Infoln(args...)
	entryPool.Put(e)
}

func (logger *Logger) Println(args ...interface{}) {
	e := NewEntry(logger)
	e.Println(args...)
	entryPool.Put(e)
}

func (logger *Logger) Warnln(args ...interface{}) {
	e := NewEntry(logger)
	e.Warnln(args...)
	entryPool.Put(e)
}

func (logger *Logger) Warningln(args ...interface{}) {
	e := NewEntry(logger)
	e.Warnln(args...)
	entryPool.Put(e)
}

func (logger *Logger) Errorln(args ...interface{}) {
	e := NewEntry(logger)
	e.Errorln(args...)
	entryPool.Put(e)
}

func (logger *Logger) Fatalln(args ...interface{}) {
	e := NewEntry(logger)
	e.Fatalln(args...)
	entryPool.Put(e)
}

func (logger *Logger) Panicln(args ...interface{}) {
	e := NewEntry(logger)
	e.Panicln(args...)
	entryPool.Put(e)
}
