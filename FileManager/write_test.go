package FileManager

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var samplePath string = "../test/SampleCodes.txt"
var sampleMessage string = "this is a sample data"

func TestOverwriteFile(t *testing.T) {
	err := OverwriteFile(samplePath, []byte(sampleMessage))
	require.NoError(t, err, "Something messed up")

	existenceChecker(t)
}

func TestOverwriteFileString(t *testing.T) {
	err := OverwriteFileString(samplePath, sampleMessage)
	require.NoError(t, err, "Something messed up")

	existenceChecker(t)
}

func TestAppendFile(t *testing.T) {
	err := AppendFile(samplePath, []byte(sampleMessage))
	require.NoError(t, err, "Something messed up")

	existenceChecker(t)
}

func TestAppendFileString(t *testing.T) {
	err := AppendFileString(samplePath, sampleMessage)
	require.NoError(t, err, "Something messed up")

	existenceChecker(t)
}

func existenceChecker(t *testing.T) {
	existence := IsFileExists(samplePath)

	t.Log("IsFileExists: ", existence)

	if existence == true {
		os.Remove(samplePath)
	}

}
