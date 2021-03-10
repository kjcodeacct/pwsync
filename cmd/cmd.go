package cmd

import (
	"fmt"
	"os/exec"
	"strings"
)

const (
	ConvertCMDType = "convert"
	LoginCMDType   = "login"
	LogoutCMDType  = "logout"
	PullCMDType    = "pull"
	PushCMDType    = "push"
)

func GetCommand(cmdName string, cfg Config, args []string) (string, error) {
	cmdFound := false
	var cmdList []string

	for _, cmd := range cfg.CmdList {
		if cmd.Name == cmdName {
			cmdFound = true
			for _, params := range cmd.CMD {
				cmdList = append(cmdList, params)
			}
		}
	}

	for _, arg := range args {
		cmdList = append(cmdList, arg)
	}

	if cmdFound == false || len(cmdList) == 0 {
		return "", fmt.Errorf("no config found for command %s", cmdName)
	}

	return strings.Join(cmdList, " "), nil
}

func RunCommand(cmd string) (string, error) {
	cmdOutput, err := exec.Command(cmd).Output()
	if err != nil {
		return "", err
	}

	return string(cmdOutput), nil
}
