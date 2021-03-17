package runtime

import (
	"os"
	"strings"

	"github.com/kjcodeacct/pwsync/platform"

	yaml "gopkg.in/yaml.v2"
)

const (
	defaultTimeout = 10
	noStdOutFile   = ""
)

type Config struct {
	Platform string    `yaml:"platform"`
	Password string    `yaml:"password,omitempty"`
	Timeout  int       `yaml:"timeout"`
	CmdList  []Command `yaml:"cmdList"`
}

type Command struct {
	Name       string   `yaml:"name"`
	CMD        []string `yaml:"cmd"`
	StdOutFile string   `yaml:"stdoutFile,omitempty"`
}

func OpenConfig(filepath string) (*Config, error) {

	buff, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	newConfig := &Config{}
	err = yaml.Unmarshal(buff, &newConfig)
	if err != nil {
		return nil, err
	}

	return newConfig, nil
}

func WriteConfig(cfg *Config, filename string) error {

	buff, err := yaml.Marshal(&cfg)
	if err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(buff)
	if err != nil {
		return err
	}

	return nil
}

func GetDefaultConfig(inputPlatform string) *Config {

	defaultCfg := &Config{}

	switch inputPlatform {
	case platform.Bitwarden:
		defaultCfg = getDefaultBwConfig()
	case platform.Lastpass:
		defaultCfg = getDefaultLpConfig()
	default:
		loginCmd := Command{
			Name: LoginCMDType,
			CMD:  []string{"exec_name_goes_here", "login"},
		}

		logoutCMD := Command{
			Name: LogoutCMDType,
			CMD:  []string{"exec_name_goes_here", "logout"},
		}

		pullCMD := Command{
			Name: PullCMDType,
			CMD:  []string{"exec_name_goes_here", "export"},
		}

		fetchCMD := Command{
			Name: FetchCMDType,
			CMD:  []string{"exec_name_goes_here", "sync"},
		}

		cmdList := []Command{loginCmd, logoutCMD, pullCMD, fetchCMD}

		defaultCfg = &Config{
			Platform: strings.Join(platform.GetSupportedPlatforms(), ","),
			Timeout:  defaultTimeout,
			CmdList:  cmdList,
		}
	}

	return defaultCfg
}

func getDefaultBwConfig() *Config {

	loginCmd := Command{
		Name: LoginCMDType,
		CMD:  []string{platform.BitwardenProcess, "login", "{PWSYNC_USERNAME}", "{PWSYNC_PASSWORD}"},
	}

	logoutCMD := Command{
		Name:       LogoutCMDType,
		CMD:        []string{platform.BitwardenProcess, "logout"},
		StdOutFile: "lastpass_export.csv",
	}

	pullCMD := Command{
		Name: PullCMDType,
		CMD:  []string{platform.BitwardenProcess, "export"},
	}

	fetchCMD := Command{
		Name: FetchCMDType,
		CMD:  []string{platform.BitwardenProcess, "sync"},
	}

	cmdList := []Command{loginCmd, logoutCMD, pullCMD, fetchCMD}

	defaultCfg := &Config{
		Platform: platform.Lastpass,
		Timeout:  defaultTimeout,
		CmdList:  cmdList,
	}

	return defaultCfg
}

func getDefaultLpConfig() *Config {

	loginCmd := Command{
		Name: LoginCMDType,
		CMD:  []string{platform.LastpassProcess, "login", "{PWSYNC_USERNAME}"},
	}

	logoutCMD := Command{
		Name: LogoutCMDType,
		CMD:  []string{platform.LastpassProcess, "logout"},
	}

	pullCMD := Command{
		Name:       PullCMDType,
		CMD:        []string{platform.LastpassProcess, "export"},
		StdOutFile: "lastpass_export.csv",
	}

	fetchCMD := Command{
		Name: FetchCMDType,
		CMD:  []string{platform.LastpassProcess, "sync"},
	}

	cmdList := []Command{loginCmd, logoutCMD, pullCMD, fetchCMD}

	defaultCfg := &Config{
		Platform: platform.Lastpass,
		Timeout:  defaultTimeout,
		CmdList:  cmdList,
	}

	return defaultCfg
}
