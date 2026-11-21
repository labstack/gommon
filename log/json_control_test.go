package log

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestJSONLogControlCharacters(t *testing.T) {
	for _, message := range []string{
		"bell: \a",
		"vertical tab: \v",
		"NUL: \x00",
		"colored output: \x1b[31mred\x1b[0m",
		"newline: \n and tab: \t",
		"quoted: \"hello\"",
	} {
		t.Run(message, func(t *testing.T) {
			var output bytes.Buffer
			logger := New("worker")
			logger.SetOutput(&output)
			logger.Info(message)

			var entry struct {
				Message string `json:"message"`
			}
			if err := json.Unmarshal(output.Bytes(), &entry); err != nil {
				t.Fatalf("invalid JSON log %q: %v", output.String(), err)
			}
			if entry.Message != message {
				t.Fatalf("message = %q, want %q", entry.Message, message)
			}
		})
	}
}

func TestJSONBodyControlCharacters(t *testing.T) {
	var output bytes.Buffer
	logger := New("worker")
	logger.SetOutput(&output)
	logger.Infoj(JSON{"message": "colored output: \x1b[31mred\x1b[0m"})
	if !json.Valid(output.Bytes()) {
		t.Fatalf("invalid JSON body log: %q", output.String())
	}
}
