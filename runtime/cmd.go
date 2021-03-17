package runtime

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"

	"github.com/creack/pty"
	"golang.org/x/term"
)

const (
	ConvertCMDType = "convert"
	InitCMDType    = "init"
	LoginCMDType   = "login"
	LogoutCMDType  = "logout"
	PullCMDType    = "pull"
	PushCMDType    = "push"
	FetchCMDType   = "fetch"
)

func GetCommand(cmdName string, cfg *Config) (string, []string, string, error) {

	cmdFound := false

	var newCMD string
	var paramList []string
	var stdoutFile string

	for _, cmd := range cfg.CmdList {
		if cmd.Name == cmdName {
			for i, arg := range cmd.CMD {
				if i == 0 {
					newCMD = arg
					cmdFound = true
				} else {
					paramList = append(paramList, arg)
				}
			}

			if cmd.StdOutFile != "" {
				stdoutFile = cmd.StdOutFile
			}
		}
	}

	for i, param := range paramList {

		envValue := checkEnv(param)
		if envValue != "" {
			paramList[i] = envValue
		}
	}

	if cmdFound == false || newCMD == "" {
		return "", nil, "", fmt.Errorf("no config found for command %s", cmdName)
	}

	return newCMD, paramList, stdoutFile, nil
}

func checkEnv(input string) string {

	var envValue string
	if strings.HasPrefix(input, "{") && strings.HasSuffix(input, "}") {
		input = strings.ReplaceAll(input, "{", "")
		input = strings.ReplaceAll(input, "}", "")
		input = strings.ReplaceAll(input, " ", "")

		envValue = os.Getenv(input)

		if envValue == "" {
			fmt.Println("env variable not found for:", input)
		}
	}

	return envValue
}

func RunCommand(cmd string, args []string, stdoutFilePath string) error {

	runCmd := exec.Command(cmd, args...)
	// start cmd with pty
	ptmx, err := pty.Start(runCmd)
	if err != nil {
		return err
	}

	// try to gracefully close pty after done
	defer func() {
		err = ptmx.Close()
		if err != nil {
			fmt.Println("error'd closing cmd:", err.Error())
		}
	}()

	// resize psuedo terminal as base terminal is modified
	sizeCh := make(chan os.Signal, 1)
	signal.Notify(sizeCh, syscall.SIGWINCH)
	go func() {
		for range sizeCh {
			pty.InheritSize(os.Stdin, ptmx)
		}
	}()

	// initialize terminal size
	sizeCh <- syscall.SIGWINCH

	// set stdin to raw mode for initial terminal state
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return err
	}

	// try to restore terminal after completion
	defer func() {
		err = term.Restore(int(os.Stdin.Fd()), oldState)
		if err != nil {
			fmt.Println(err)
		}
	}()

	// Copy stdin to the pty and the pty to stdout.
	go func() {
		io.Copy(ptmx, os.Stdin)
	}()

	if stdoutFilePath != "" {
		stdoutFile, err := os.Create(stdoutFilePath)
		if err != nil {
			return err
		}
		defer stdoutFile.Close()

		io.Copy(stdoutFile, ptmx)
	} else {
		io.Copy(os.Stdout, ptmx)
	}

	return nil
}
