package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   LoginCMDType,
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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

	cmd, params, err := GetCommand(LoginCMDType, userCfg)
	if err != nil {
		return err
	}

	params = append(params, args...)

	err = RunCommand(cmd, params)
	if err != nil {
		return err
	}

	return nil
}
