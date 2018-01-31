package services

import (
  "io/ioutil"
  "testing"
  "github.com/stretchr/testify/assert"
)

var exampleDiscoveryXmlPath = "./../test/fixtures/discovery.xml"

func TestCreateWopiDiscovery(t *testing.T) {
    discovery := CreateWopiDiscovery()

    assert.NotNil(t, discovery)
}

func TestParseDiscoveryXml(t *testing.T) {
  data, err := ioutil.ReadFile(exampleDiscoveryXmlPath)

  discovery, err := ParseDiscoveryXml(data)

  assert.Nil(t, err)
  assert.NotNil(t, discovery)
  assert.Equal(t, len(discovery.NetZones), 2)
}