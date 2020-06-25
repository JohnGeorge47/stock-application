package configmanager

import (
	"encoding/json"
	"errors"
	"os"
)

// ApplicationConfig is the configuration for the ApplicationConfig
type ApplicationConfig struct {
	ProcessName string    `json:"process_name,omitempty"`
	Origins     []string  `json:"origins"`
	RedisConfig RdbConfig `json:"redis_config"`
	SqlConf     MySqlConf `json:"sql_conf"`
	ProcessTime int			`json:"process_time"`
}

type MySqlConf struct {
	DbUser string `json:"db_user"`
	DbPwd  string `json:"db_pwd"`
	DbPort string `json:"db_port"`
	DbHost string `json:"db_host"`
	DbName string `json:"db_name"`
}

type RdbConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
	Db   int    `json:"db"`
}

// Config stores the configuration
var Config *ApplicationConfig
var configFile *string

// LoadConfiguration loads configuration from file
// decodes the json config file into an instance of application config
// if the decoded config is valid it is set as config
func LoadConfiguration() error {
	if configFile == nil {
		return errors.New("config not initialized")
	}
	config := new(ApplicationConfig)
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
func GetConfig() ApplicationConfig {
	return *Config
}

// InitConfig will initialize app config with config file name
func InitConfig(config *string) error {
	configFile = config
	return LoadConfiguration()
}
