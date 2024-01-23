package logx

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
)

func TestColor(t *testing.T) {
	result := GREEN("This is a test!")
	if result != "\x1b[32mThis is a test!\x1b[0m" {
		t.Errorf("%s should be a green log that says 'This is a test!'", result)
	}
}

func TestLogWithLevelDebug(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()
	LogWithLevel(Debug, "This is a debug log")
	if !strings.Contains(buf.String(), "😎  \x1b[36mDEBUG\x1b[0m - This is a debug log\n") {
		t.Errorf("Improperly formatted log: %s", buf.String())
	}
	t.Log(buf.String())
}

func TestLogWithLevelInfo(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()
	LogWithLevel(Info, "This is informative")
	if !strings.Contains(buf.String(), " ℹ️  \x1b[32mINFO\x1b[0m - This is informative\n") {
		t.Errorf("Improperly formatted log: %s", buf.String())
	}
	t.Log(buf.String())
}

func TestLogWithLevelNotice(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()
	LogWithLevel(Notice, "This is a notice level log")
	if !strings.Contains(buf.String(), "📌  \x1b[33mNOTICE\x1b[0m - This is a notice level log\n") {
		t.Errorf("Improperly formatted log: %s", buf.String())
	}
	t.Log(buf.String())
}

func TestLogWithLevelWarning(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()
	LogWithLevel(Warning, "This is a warning level log")
	if !strings.Contains(buf.String(), "⚠️  \x1b[35mWARNING\x1b[0m - This is a warning level log\n") {
		t.Errorf("Improperly formatted log: %s", buf.String())
	}
	t.Log(buf.String())
}

func TestLogWithLevelError(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()
	LogWithLevel(Error, "This is a error level log")
	if !strings.Contains(buf.String(), "💩  \x1b[31mERROR\x1b[0m - This is a error level log\n") {
		t.Errorf("Improperly formatted log: %s", buf.String())
	}
	t.Log(buf.String())
}

func TestLogWithLevelCritical(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()
	LogWithLevel(Critical, "This is a critical level log")
	if !strings.Contains(buf.String(), "😱  \x1b[31mCRITICAL\x1b[0m - This is a critical level log\n") {
		t.Errorf("Improperly formatted log: %s", buf.String())
	}
	t.Log(buf.String())
}

func TestLogWithLevelAlert(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()
	LogWithLevel(Alert, "This is a alert level log")
	if !strings.Contains(buf.String(), "☠️  \x1b[31mALERT\x1b[0m - This is a alert level log\n") {
		t.Errorf("Improperly formatted log: %s", buf.String())
	}
	t.Log(buf.String())
}

func TestLogWithLevelEmergency(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()
	LogWithLevel(Emergency, "This is a emergency level log!!!")
	if !strings.Contains(buf.String(), "☢️  \x1b[41mEMERGENCY\x1b[0m - This is a emergency level log!!!\n") {
		t.Errorf("Improperly formatted log: %s", buf.String())
	}
	t.Log(buf.String())
}