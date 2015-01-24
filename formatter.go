package logrus

import (
	"bytes"
)

// The Formatter interface is used to implement a custom Formatter. It takes an
// `Entry`. It exposes all the fields, including the default ones:
//
// * `entry.Data["msg"]`. The message passed from Info, Warn, Error ..
// * `entry.Data["time"]`. The timestamp.
// * `entry.Data["level"]. The level the entry was logged at.
//
// Any additional fields added with `WithField` or `WithFields` are also in
// `entry.Data`. Format is expected to return an array of bytes which are then
// logged to `logger.Out`.
type Formatter interface {
	Format(*Entry, *bytes.Buffer) error
}

// This is to not silently overwrite `time`, `msg` and `level` fields when
// dumping it. If this code wasn't there doing:
//
//  logrus.WithField("level", 1).Info("hello")
//
// Would just silently drop the user provided level. Instead with this code
// it'll logged as:
//
//  {"level": "info", "fields.level": 1, "msg": "hello", "time": "..."}
//
// It's not exported because it's still using Data in an opinionated way. It's to
// avoid code duplication between the two default formatters.
func prefixFieldClashes(data Fields) Fields {
	cloned := false

	for k, r := range prefixKeys {
		_, ok := data[k]
		if ok {
			if !cloned {
				data = cloneData(data)
				cloned = true
			}
			data[r] = data[k]
		}
	}

	return data
}

var prefixKeys = map[string]string{
	"time":  "fields.time",
	"msg":   "fields.msg",
	"level": "fields.level",
}
