package cmd

import (
	"fmt"
	"os"

	"github.com/kjcodeacct/pwsync/runtime"
	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   runtime.LoginCMDType,
	Short: "Login to your password manager.",
	Long:  `Login to the desired password platform, executes the 'login' command desginated by the password manager.`,
	Run: func(cmd *cobra.Command, args []string) {

		err := login(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}

func login(args []string) error {

	cmd, params, stdoutFile, err := runtime.GetCommand(runtime.LoginCMDType, userCfg)
	if err != nil {
		return err
	}

	params = append(params, args...)

	err = runtime.RunCommand(cmd, params, stdoutFile)
	if err != nil {
		return err
	}

	return nil
}
