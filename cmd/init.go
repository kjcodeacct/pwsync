package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
)

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
}

func initPwSync() error {

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	fileExists := true
	if _, err := os.Stat(filepath.Join(cwd, cfgFile)); os.IsNotExist(err) {
		fileExists = false
	}

	if fileExists {
		return fmt.Errorf("config already exists: %s", cfgFile)
	}

	defaultCfg := GetDefaultConfig()
	err = WriteConfig(defaultCfg, filepath.Join(cwd, cfgFile))
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
