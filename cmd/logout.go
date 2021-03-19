package cmd

import (
	"fmt"
	"os"

	"github.com/kjcodeacct/pwsync/runtime"
	"github.com/spf13/cobra"
)

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   runtime.LogoutCMDType,
	Short: "Logout of your password manager.",
	Long:  `Logout of the desired password platform, executes the 'logout' command desginated by the password manager.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := logout(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}

func logout(args []string) error {
	cmd, params, stdoutFile, err := runtime.GetCommand(runtime.LogoutCMDType, userCfg)
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
