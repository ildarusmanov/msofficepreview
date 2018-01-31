package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	ServerHost      string `yaml:"server_host"`
	ServerAddr      string `yaml:"server_addr"`
	DiscoveryXmlUrl string `yaml:"discovery_xml_url"`
	StorageRoot     string `yaml:"storage_root"`
	TokenLifetime   int64  `yaml:"token_lifetime"`
}

/**
 * Create new config object
 * @return *Config
 */
func CreateNewConfig() *Config {
	return &Config{}
}

/**
 * Load data from file
 * @param string
 * @return *Config
 */
func LoadConfigYAML(configFilePath string) (*Config, error) {
	config := CreateNewConfig()

	configFileData, err := ioutil.ReadFile(configFilePath)

	if err != nil {
		return nil, err
	}

	if err = yaml.Unmarshal([]byte(configFileData), config); err != nil {
		return nil, err
	}

	return config, nil
}
