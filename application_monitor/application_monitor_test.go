package application_monitor

import (
	"os"
	"os/user"
	"reflect"
	"testing"
)

func TestGetSampleConfig(t *testing.T) {
	sampleConfig := getString()
	vmType := reflect.TypeOf(sampleConfig)
	if k := vmType.Kind(); k != reflect.String {
		t.Error("Wrong type received!\nKind found: ", vmType.Kind())
	}
}

func TestGetDescription(t *testing.T) {
	description := getString()
	vmType := reflect.TypeOf(description)
	if k := vmType.Kind(); k != reflect.String {
		t.Error("Wrong type received!\nKind found: ", vmType.Kind())
	}
}

func TestGetMetricsFromFile(t *testing.T) {
	// Retrieve the current user
	var curUser, err = user.Current()

	// Test fails when any error occurs or when the current user could not be retrieved
	if err != nil || curUser == nil {
		t.Error("No user found!\n",err)
	}

	// Get the current users' home directory
	var PATH = curUser.HomeDir

	// Get the file from the users' desktop
	var FILE = PATH + `\Desktop\monitor.log.txt`

	// Attempt to open the specified file
	file, err := os.Open(FILE)

	// Test fails when the file could not be opened for any reason or if any other error occurs
	if file == nil || err != nil {
		t.Error("Failed to open file!\n",err)
	}
}

func TestWriteToInfluxDatabase(t *testing.T) {

}

func getString() string {
	return `Sample Config`
}