package FileManager

import (
	"os"
	"testing"
)

var samplePath string = "../test/SampleCodes.txt"
var sampleMessage string = "this is a sample data"

func TestOverwriteFile(t *testing.T) {
	t.Log(OverwriteFile(samplePath, []byte(sampleMessage)))

	existenceChecker(t)
}

func TestOverwriteFileString(t *testing.T) {
	t.Log(OverwriteFileString(samplePath, sampleMessage))

	existenceChecker(t)
}

func TestAppendFile(t *testing.T) {
	t.Log(AppendFile(samplePath, []byte(sampleMessage)))

	existenceChecker(t)
}

func TestAppendFileString(t *testing.T) {
	t.Log(AppendFileString(samplePath, sampleMessage))

	existenceChecker(t)
}

func existenceChecker(t *testing.T) {
	existence := IsFileExists(samplePath)

	t.Log("IsFileExists: ", existence)

	if existence == true {
		os.Remove(samplePath)
	}

}
