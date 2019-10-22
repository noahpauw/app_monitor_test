package application_monitor

import (
	"fmt"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"

	"io/ioutil"
	"os"
	user2 "os/user"
	"time"
)

// TODO: Setup communication with logging application file
type ApplicationMonitor struct {
	MsrCPU  bool   `toml:"msrCPU"`
	MsrVM   bool   `toml:"msrVM"`
	appName string `toml:"appName"`
}

var sampleConfig = `
	## Whether to include CPU measurements from the file
	msrCPU = true
	## Whether to include IO measurements from the file
	msrVM = true
	## The name of the application you want to monitor
	appName = "Monitor"
`

const (
	CPU = 1
	VIRTUAL_MEMORY = 2
)

func (_ *ApplicationMonitor) SampleConfig() string {
	return sampleConfig
}

var description = `
	# This plugin contacts a custom performance counter that measures metrics in any system.
	# This application is very flexible and can monitor anything on the application-layer
`

func (_ *ApplicationMonitor) Description() string {
	return description
}

var user, _ = user2.Current()
var FILE = user.HomeDir + "/Desktop/monitor.log.txt"

func (i *ApplicationMonitor) Gather(accumulator telegraf.Accumulator) error {
	now := time.Now()
	_ = now

	// TODO: Remove this as it should read from a file
	var cpu uint64 = cpu2.Percent(5 * time.Second, true)

	// Gather some CPU and IO stats
	if i.MsrCPU {
		// Read CPU from logging file
		readFile(CPU)
		accumulator.addFields("cpuUsage",map[string]interface{}{"cpu": cpu}, nil)
	}

	// TODO: Remove this as it should read from a file
	var virtMem uint64 = mem.VirtualMemory()

	// Only read IO if specified
	if i.MsrVM {
		// Read IO from logging file
		readFile(VIRTUAL_MEMORY)
		accumulator.addFields("vmUsage",map[string]interface{}{"vm": virtMem}, nil)
	}

	fmt.Println(i.appName)
	return nil
}

/**
 * readFile(int)
 * Reads a file from the desktop and converts the data
 * Requires an integer value to determine which data to gather
 * 1 = CPU
 * 2 = Virtual Memory
 */
func ReadFile(measure int) {
	dat, err := ioutil.ReadFile(FILE)
	check(err)

	// Solve annoying Go error
	_ = dat

	// Open the file if it exists
	f, err := os.Open(FILE)

	// Check if any error occur
	check(err)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)

	// Check if any error occur
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))
}

/*
	* The init function is used to register a plugin to Telegraf.
	* The plugin uses the Add function to add itself as a valid plugin
*/
func init() {
	inputs.Add("application_monitor", func() telegraf.Input {
		return &ApplicationMonitor{
			MsrCPU:  true,
			MsrVMr:   true,
			appName: "application_monitor",
		}
	})
}

// Checks if an error occurs and ends the program if anything happens
func Check(err error) {
	if err != nil {
		panic(err)
	}
}
