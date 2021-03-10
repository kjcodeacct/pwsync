package cmd

import (
	"os"
	"strings"

	"github.com/kjcodeacct/pwsync/platform"
	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Platform string    `yaml:"platform"`
	CmdList  []Command `yaml:"cmdList"`
}

type Command struct {
	Name string   `yaml:"name"`
	CMD  []string `yaml:"cmd"`
}

func Open(filepath string) (*Config, error) {

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

func GetDefaultConfig() *Config {

	loginCmd := Command{
		Name: LoginCMDType,
	}

	logoutCMD := Command{
		Name: LogoutCMDType,
	}

	pullCMD := Command{
		Name: PullCMDType,
	}

	pushCmd := Command{
		Name: PushCMDType,
	}

	cmdList := []Command{loginCmd, logoutCMD, pullCMD, pushCmd}

	defaultCfg := Config{
		Platform: strings.Join(platform.GetSupportedPlatforms(), ","),
		CmdList:  cmdList,
	}

	return &defaultCfg
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
