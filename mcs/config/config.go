package config

import (
	"os"
	"path/filepath"

	"go-mcs-sdk/mcs/api/common/logs"

	"github.com/BurntSushi/toml"
)

type Configuration struct {
	Apikey      string `toml:"apikey"`
	AccessToken string `toml:"access_token"`
	Network     string `toml:"network"`
	File2Upload string `toml:"file_to_upload"`
	RpcUrl      string `toml:"rpc_url"`
	PrivateKey  string `toml:"private_key"`
}

var config *Configuration

func InitConfig() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		logs.GetLogger().Fatal("Cannot get home directory.")
	}

	configFile := filepath.Join(homedir, ".swan/mcs_sdk", "config.toml")

	if metaData, err := toml.DecodeFile(configFile, &config); err != nil {
		logs.GetLogger().Fatal("error:", err)
	} else {
		if !requiredFieldsAreGiven(metaData) {
			logs.GetLogger().Fatal("required fields not given")
		}
	}
}

func GetConfig() *Configuration {
	if config == nil {
		InitConfig()
	}
	return config
}

func requiredFieldsAreGiven(metaData toml.MetaData) bool {
	requiredFields := [][]string{
		{"apikey"},
		{"access_token"},
		{"network"},
		{"file_to_upload"},
	}

	for _, v := range requiredFields {
		if !metaData.IsDefined(v...) {
			logs.GetLogger().Fatal("required fields ", v)
		}
	}

	return true
}
