package FileManager

import (
	"fmt"
	"os"
	"testing"
)

var samplePath string = "../test/SampleCodes.md"
var sampleMessage string = "this is a sample data"

func TestOverwriteFile(t *testing.T) {
	t.Log(OverwriteFile(samplePath, []byte(sampleMessage)))

	_, err := os.Stat(samplePath)

	if err != nil {
		t.Log("Oh fucked: %w", err)
	} else {
		t.Log("Fine, file exist")
		os.Remove(samplePath)
	}
}

func TestOverwriteFileString(t *testing.T) {
	t.Log(OverwriteFileString(samplePath, sampleMessage))

	_, err := os.Stat(samplePath)

	if err != nil {
		t.Log("Oh fucked: %w", err)
	} else {
		t.Log("Fine, file exist")
		os.Remove(samplePath)
	}
}

func TestAppendFile(t *testing.T) {
	t.Log(AppendFile(samplePath, []byte(sampleMessage)))

	_, err := os.Stat(samplePath)

	if err != nil {
		t.Log("Oh fucked: %w", err)
	} else {
		t.Log("Fine, file exist")
		os.Remove(samplePath)
	}
}

func TestAppendFileString(t *testing.T) {
	t.Log(AppendFileString(samplePath, sampleMessage))

	_, err := os.Stat(samplePath)

	if err != nil {
		t.Log("Oh fucked: %w", err)
	} else {
		t.Log("Fine, file exist")
		os.Remove(samplePath)
	}

}
