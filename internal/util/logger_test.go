package util

import (
	"bytes"
	"log"
	"testing"
)

func TestLogger(t *testing.T) {
	var buf bytes.Buffer
	logger := log.New(&buf, "INFO: ", log.Lshortfile)

	logger.Println("This is a test log")
	if !bytes.Contains(buf.Bytes(), []byte("This is a test log")) {
		t.Errorf("expected log message to contain 'This is a test log', got %s", buf.String())
	}
}
