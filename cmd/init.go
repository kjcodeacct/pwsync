package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/kjcodeacct/pwsync/platform"
	"github.com/kjcodeacct/pwsync/runtime"
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
)

var initPlatform string

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the working directory.",
	Long:  `Initialize the working directory for pwsync, and create the configuration file.`,
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
	initCmd.PersistentFlags().StringVar(&initPlatform, "platform", "",
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

	defaultCfg := runtime.GetDefaultConfig(initPlatform)

	err := runtime.WriteConfig(defaultCfg, filepath.Join(currentDir, cfgFile))
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
