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
	Short: "Fetch your latest password vault.",
	Long:  `Fetches the latest password vault from desired password platform, executes the 'sync/fetch/update' command designated by the password manager.`,
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
