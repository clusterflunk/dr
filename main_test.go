package main

import (
	"testing"
)

func TestConfigReadWrite(t *testing.T) {
	configFilePath := "dr.json.test"
	config := NewDrConfig()
	cli := NewDrCli(config)

	// write an empty config
	cli.writeConfigFile(configFilePath)

	run1 := DrRun{I: true}
	run2 := DrRun{I: true, Command: "blah"}
	run3 := DrRun{I: false}
	run4 := DrRun{V: "this:that"}
	run5 := DrRun{I: true}
	cli.config.Runs = []DrRun{run1, run2, run3, run4, run5}
	cli.writeConfigFile(configFilePath)
}
