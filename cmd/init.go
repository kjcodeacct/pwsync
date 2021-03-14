package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/kjcodeacct/pwsync/platform"
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
)

var initPlatform string

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		err := initPwSync()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	rootCmd.PersistentFlags().StringVar(&initPlatform, "platform", "",
		fmt.Sprintf("platform to create a default cfg (%s)",
			strings.Join(platform.GetSupportedPlatforms(), ",")))

}

func initPwSync() error {
	fileExists := true
	if _, err := os.Stat(filepath.Join(currentDir, cfgFile)); os.IsNotExist(err) {
		fileExists = false
	}

	if fileExists {
		return fmt.Errorf("config already exists: %s", cfgFile)
	}

	defaultCfg := GetDefaultConfig(initPlatform)

	err := WriteConfig(defaultCfg, filepath.Join(currentDir, cfgFile))
	if err != nil {
		return err
	}

	fmt.Println("initialized pwsync:", cfgFile)
	fmt.Println("attempting to open:", cfgFile)
	err = open.Run(cfgFile)
	if err != nil {
		return err
	}

	return nil
}
