package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/kjcodeacct/pwsync/files"
	"github.com/kjcodeacct/pwsync/platform"
	"github.com/kjcodeacct/pwsync/runtime"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var cleanupPullFiles bool

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   runtime.PullCMDType,
	Short: "Export and convert your password vault into a keepass database.",
	Long:  `Exports the latest password vault into a CSV, then converts it to a keepass database`,
	Run: func(cmd *cobra.Command, args []string) {

		err := pull(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)
	pullCmd.PersistentFlags().BoolVar(&cleanupPullFiles, "cleanup", true,
		"auto cleanup pulled files")
}

func pull(args []string) error {
	cmd, params, stdoutFile, err := runtime.GetCommand(runtime.PullCMDType, userCfg)
	if err != nil {
		return err
	}

	params = append(params, args...)

	fileCh := files.ListenForType(currentDir, files.CSVExtension)

	err = runtime.RunCommand(cmd, params, stdoutFile)
	if err != nil {
		return err
	}

	var platformExportFile string

	select {
	case filepath := <-fileCh:
		platformExportFile = filepath
	case <-time.After(time.Second * time.Duration(userCfg.Timeout)):
		fmt.Printf("failed to find exported password db after 10 seconds, please check %s export file: %s\n",
			userCfg.Platform, platformExportFile)
		return nil
	}

	defer func() {
		if cleanupPullFiles && platformExportFile != "" {
			err = files.Cleanup(platformExportFile, 1)
			if err != nil {
				fmt.Println("error cleaning up file", err)
			}

			fmt.Println("cleaned up file:", platformExportFile)
		}
	}()

	kpExport, err := platform.ConvertCSV(userCfg.Platform, platformExportFile)
	if err != nil {
		return err
	}

	if len(kpExport.KeepassGroup.Entries) > 0 || len(kpExport.KeepassGroup.Groups) > 0 {
		if userCfg.Password == "" {

			prompt := promptui.Prompt{
				Label: "Backup Password",
				Mask:  '*',
			}

			password, err := prompt.Run()
			if err != nil {
				return err
			}

			userCfg.Password = password
		}

		newKpFilePath, err := kpExport.Write(userCfg.Password)
		if err != nil {
			return err
		}

		fmt.Println("pulled backup to:", newKpFilePath)
	} else {
		fmt.Printf("no data backed up to database, please check %s export file: %s\n",
			userCfg.Platform, platformExportFile)
		cleanupPullFiles = false
	}

	return nil
}
