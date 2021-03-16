package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string

var userCfg *Config

var currentDir string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pwsync",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	var err error
	// Find current working directory.
	currentDir, err = os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", ".pwsync.yaml", "Config file for pwsync")
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {

		// Search config in home directory with name ".pwsync" (without extension).
		viper.AddConfigPath(currentDir)
		viper.SetConfigName(".pwsync")
	}

	viper.AutomaticEnv() // read in environment variables that match

	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("no command provided")
	}

	if args[0] != InitCMDType {
		var err error
		userCfg, err = OpenConfig(cfgFile)
		if err != nil {
			fmt.Println(err)
			fmt.Println("please check pwsync config, or run 'pwsync init'")
			os.Exit(1)
		}
	}

}
