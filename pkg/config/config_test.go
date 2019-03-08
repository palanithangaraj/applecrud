package config

import (
	"bytes"
	"os"
	"strings"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestGetConfigFromENV(t *testing.T) {
	os.Setenv("PORT", "3000")
	os.Setenv("PRELOG", "TNG - applecrud - main.go: ")
	os.Setenv("LOG_LEVEL", "trace")

	//give it a path that should not work and force it to go to the env variables
	_, err := GetConfig("./failure.json")

	if err != nil {
		t.Error("unable to read ENV variables.", err)
	}
}

func TestGetConfigFromENV2(t *testing.T) {
	os.Setenv("PORT", "3000")
	//setting to "" instead of not including so when running the package of tests, it does not use the EVN variable from the above test.
	os.Setenv("PRELOG", "")
	os.Setenv("LOG_LEVEL", "trace")

	//give it a path that should not work and force it to go to the env variables
	_, err := GetConfig("./failure.json")

	//there should be an error as one of the ENV variables is missing.
	if err == nil {
		t.Error("expected an error on the configuration file for a missing ENV variable. error was nil")
	}
}

func TestDecodeURL(t *testing.T) {
	_, err := decodeURL("bW9uZ29kYjovL2FkbWluOmF2ZmtBV0NMRGFBTWN2Tm1RWVFzcnhTNUAxMC4zLjAuMjQzOjI3MDE3")
	if err != nil {
		t.Error("unable to decode encoded URL", err)
	}
}

func TestLevel_Info_AtLevel(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	log.SetLevel(log.InfoLevel)

	log.Info("123")
	if !strings.Contains(buf.String(), "123") {
		t.Errorf("expected 123 output, but got: %v", buf.String())
	}
}

func TestLevel_Error_BelowLevel(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	log.SetLevel(log.PanicLevel)

	log.Error("123")
	if buf.String() != "" {
		t.Errorf("expected empty output, but got: %v", buf.String())
	}
}

func TestLevel_Info_AboveLevel(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	log.SetLevel(log.ErrorLevel)

	log.Error("123")
	if !strings.Contains(buf.String(), "123") {
		t.Errorf("expected 123 output, but got: %v", buf.String())
	}
}

func TestPrint(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	log.SetLevel(log.TraceLevel)

	c := Config{
		LogLevel: log.TraceLevel,
	}

	(&c).Print()
	if !strings.Contains(buf.String(), "log-level: trace") {
		t.Errorf("expected trace log-level output, but got: %v", buf.String())
	}
}

func TestPrint_BelowLevel(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	log.SetLevel(log.InfoLevel)

	c := Config{
		LogLevel: log.InfoLevel,
	}

	(&c).Print()
	if buf.String() != "" {
		t.Errorf("expected empty output, but got: %v", buf.String())
	}
}
