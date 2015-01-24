package logrus

import (
	"bytes"
	"fmt"
	"regexp"
	"sort"
	"strings"
	"time"
)

const (
	nocolor = 0
	red     = 31
	green   = 32
	yellow  = 33
	blue    = 34
)

var (
	baseTimestamp time.Time
	isTerminal    bool
	noQuoteNeeded *regexp.Regexp
)

func init() {
	baseTimestamp = time.Now()
	isTerminal = IsTerminal()
}

func miniTS() int {
	return int(time.Since(baseTimestamp) / time.Second)
}

type TextFormatter struct {
	// Set to true to bypass checking for a TTY before outputting colors.
	ForceColors   bool
	DisableColors bool
	// Set to true to disable timestamp logging (useful when the output
	// is redirected to a logging system already adding a timestamp)
	DisableTimestamp bool
}

func (f *TextFormatter) Format(entry *Entry, out *bytes.Buffer) error {

	var keys []string
	for k := range entry.Data {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	data := prefixFieldClashes(entry.Data)

	isColored := (f.ForceColors || isTerminal) && !f.DisableColors

	if isColored {
		printColored(out, entry, keys, data)
	} else {
		if !f.DisableTimestamp {
			f.appendKeyValue(out, "time", entry.Time.Format(time.RFC3339))
		}
		f.appendKeyValue(out, "level", entry.Level.String())
		f.appendKeyValue(out, "msg", entry.Message)
		for _, key := range keys {
			f.appendKeyValue(out, key, data[key])
		}
	}

	out.WriteByte('\n')
	return nil
}

func printColored(b *bytes.Buffer, entry *Entry, keys []string, data Fields) {
	var levelColor int
	switch entry.Level {
	case WarnLevel:
		levelColor = yellow
	case ErrorLevel, FatalLevel, PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}

	levelText := strings.ToUpper(entry.Level.String())[0:4]

	fmt.Fprintf(b, "\x1b[%dm%s\x1b[0m[%04d] %-44s ", levelColor, levelText, miniTS(), entry.Message)
	for _, k := range keys {
		v := data[k]
		fmt.Fprintf(b, " \x1b[%dm%s\x1b[0m=%v", levelColor, k, v)
	}
}

func needsQuoting(text string) bool {
	for _, ch := range text {
		if !((ch >= 'a' && ch <= 'z') ||
			(ch >= 'A' && ch <= 'Z') ||
			(ch >= '0' && ch < '9') ||
			ch == '-' || ch == '.') {
			return false
		}
	}
	return true
}

func (f *TextFormatter) appendKeyValue(b *bytes.Buffer, key, value interface{}) {
	switch value.(type) {
	case string:
		if needsQuoting(value.(string)) {
			fmt.Fprintf(b, "%v=%s ", key, value)
		} else {
			fmt.Fprintf(b, "%v=%q ", key, value)
		}
	case error:
		if needsQuoting(value.(error).Error()) {
			fmt.Fprintf(b, "%v=%s ", key, value)
		} else {
			fmt.Fprintf(b, "%v=%q ", key, value)
		}
	default:
		fmt.Fprintf(b, "%v=%v ", key, value)
	}
}
