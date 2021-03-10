package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   PullCMDType,
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pull called")
		err := pull(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)
}

func pull(args []string) error {

	cmd, params, err := GetCommand(PullCMDType, userCfg)
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