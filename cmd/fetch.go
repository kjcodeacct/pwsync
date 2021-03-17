package cmd

import (
	"fmt"
	"os"

	"github.com/kjcodeacct/pwsync/runtime"
	"github.com/spf13/cobra"
)

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   runtime.FetchCMDType,
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		err := fetch(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)
}

func fetch(args []string) error {

	cmd, params, stdoutFile, err := runtime.GetCommand(runtime.FetchCMDType, userCfg)
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
