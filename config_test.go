package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

var configYamlPath = "./test/fixtures/config.yml"

func TestLoadConfigYAML(t *testing.T) {
    config, err := LoadConfigYAML(configYamlPath)

    assert := assert.New(t)
    assert.Nil(err)
    assert.NotNil(config)
    assert.NotNil(config.DiscoveryXmlUrl)
    assert.NotNil(config.StorageRoot)
    assert.NotNil(config.TokenLifetime)
    assert.NotNil(config.ServerAddr)
    assert.NotNil(config.ServerHost)
}
