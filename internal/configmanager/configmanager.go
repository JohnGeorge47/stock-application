package configmanager

import (
	"encoding/json"
	"errors"
	"os"
)

// ApplicationConfig is the configuration for the ApplicationConfig
type ApplicatonConfig struct {
	ProcessName string   `json:"process_name,omitempty"`
	Origins     []string `json:"origins"`
}

// Config stores the configuration
var Config *ApplicatonConfig
var configFile *string

// LoadConfiguration loads configuration from file
// decodes the json config file into an instance of application config
// if the decoded config is valid it is set as config
func LoadConfiguration() error {
	if configFile == nil {
		return errors.New("config not initialized")
	}
	config := new(ApplicatonConfig)
	file, err := os.Open(*configFile)
	if err != nil {
		return err
	}
	if err := json.NewDecoder(file).Decode(config); err != nil {
		return err
	}
	Config = config
	return nil
}

// GetConfig returns a copy of init-ed application config instance
// if not already initialized it is initialized with "." as config path
func GetConfig() ApplicatonConfig {
	return *Config
}

// InitConfig will initialize app config with config file name
func InitConfig(config *string) error {
	configFile = config
	return LoadConfiguration()
}
