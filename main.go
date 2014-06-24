package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
)

type DrRun struct {
	Rm      bool
	I       bool
	T       bool
	V       string
	P       string
	Image   string
	Command string
}

type DrConfig struct {
	Runs []DrRun
}

func NewDrConfig() *DrConfig {
	return &DrConfig{}
}

type DrCli struct {
	configFileName string
	config         *DrConfig
}

func (cli *DrCli) readConfigFile(fileName string) {
	ba, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Had trouble reading your %s file", fileName)
	}
	err = json.Unmarshal(ba, cli.config)
	if err != nil {
		fmt.Printf("Had trouble parsing your %s file, err:", fileName, err)
	}
}

func (cli *DrCli) writeConfigFile(fileName string) {
	ba, err := json.MarshalIndent(cli.config, "", "\t")
	if err != nil {
		fmt.Printf("Had trouble reading your %s file", fileName)
	}
	err = ioutil.WriteFile(fileName, ba, 0744)
	if err != nil {
		fmt.Printf("Had trouble writing your %s file", fileName)
	}
}

func (cli *DrCli) parseCommand(args []string) error {
	if len(args) > 0 {
		return cli.RunCommand(args)
	}
	return cli.HelpCommand(args)
}

func (cli *DrCli) RunCommand(args []string) error {
	var runArgs []string

	runArgs = append(runArgs, "run")
	runArgs = append(runArgs, args...)

	cmd := exec.Command("docker", runArgs...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	return err
}

func (cli *DrCli) HelpCommand(args []string) error {
	var err error
	fmt.Println("Call any `docker run` command with the --name flag to store that run.")
	return err
}

func NewDrCli(c *DrConfig) *DrCli {
	return &DrCli{config: c}
}

func main() {
	var configFilePath string
	cwdPath := "dr.json"
	homePath := path.Join(os.Getenv("HOME"), ".dr.json")

	config := NewDrConfig()
	cli := NewDrCli(config)

	if _, err := os.Stat(cwdPath); err == nil {
		configFilePath = cwdPath
	} else if _, err := os.Stat(homePath); err == nil {
		configFilePath = homePath
	} else {
		cli.writeConfigFile(homePath)
		configFilePath = homePath
	}

	cli.readConfigFile(configFilePath)
	err := cli.parseCommand(os.Args[1:])
	if err != nil {
		fmt.Println(err)
	}
}
