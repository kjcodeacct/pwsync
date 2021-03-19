package cmd

import (
	"fmt"
	"os"

	"github.com/kjcodeacct/pwsync/runtime"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string

var userCfg *runtime.Config

var currentDir string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pwsync",
	Short: "Backup and store your desired password vault into a keepass database.",
	Long: `Pwsync is a convenient password backup tool to help with the following:

* Backup proprietary password vaults into an encrypted keepass database.

If you work with multiple password systems, or want backups of system critical passwords, this is for you.

For more information on the keepass database format please visit:
	https://keepass.info/index.html
`,
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

	if args[0] != runtime.InitCMDType {
		var err error
		userCfg, err = runtime.OpenConfig(cfgFile)
		if err != nil {
			fmt.Println(err)
			fmt.Println("please check pwsync config, or run 'pwsync init'")
			os.Exit(1)
		}
	}

}
