package cmd

import (
	"fmt"
	"os"
	"strings"

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
		fmt.Println("login called")
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

	cmd, err := GetCommand(LoginCMDType, userCfg, args)
	if err != nil {
		return err
	}

	output, err := RunCommand(cmd)
	if err != nil {
		return err
	}

	printLines := strings.Split(output, "\n")

	for _, printLine := range printLines {
		fmt.Println(printLine)
	}

	return nil
}
