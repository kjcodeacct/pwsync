package cmd

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"

	"github.com/kr/pty"
	"golang.org/x/term"
)

const (
	ConvertCMDType = "convert"
	InitCMDType    = "init"
	LoginCMDType   = "login"
	LogoutCMDType  = "logout"
	PullCMDType    = "pull"
	PushCMDType    = "push"
)

func GetCommand(cmdName string, cfg *Config) (string, []string, error) {
	cmdFound := false

	var newCMD string
	var paramList []string

	for _, cmd := range cfg.CmdList {
		if cmd.Name == cmdName {
			for i, arg := range cmd.CMD {
				if i == 0 {
					newCMD = arg
					cmdFound = true
				} else {
					tempList := strings.Split(arg, " ")
					paramList = append(paramList, tempList...)
				}
			}
		}
	}

	if cmdFound == false || newCMD == "" {
		return "", nil, fmt.Errorf("no config found for command %s", cmdName)
	}

	return newCMD, paramList, nil
}

func RunCommand(cmd string, args []string) error {

	c := exec.Command(cmd, args...)
	// start cmd with pty
	ptmx, err := pty.Start(c)
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

	io.Copy(os.Stdout, ptmx)

	return nil
}
