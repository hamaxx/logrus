package logrus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

type JSONFormatter struct{}

func (f *JSONFormatter) Format(entry *Entry, out *bytes.Buffer) error {
	data := prefixFieldClashes(entry.Data)

	data["time"] = entry.Time.Format(time.RFC3339)
	data["msg"] = entry.Message
	data["level"] = entry.Level.String()

	err := json.NewEncoder(out).Encode(data)
	if err != nil {
		return fmt.Errorf("Failed to marshal fields to JSON, %v", err)
	}

	// Cleanup special data
	for k := range prefixKeys {
		if entry.Data[k] == data[k] {
			delete(entry.Data, k)
		}
	}

	return nil
}

func cloneData(data Fields) Fields {
	newData := make(Fields, len(data))
	for k, v := range data {
		newData[k] = v
	}
	return newData
}
