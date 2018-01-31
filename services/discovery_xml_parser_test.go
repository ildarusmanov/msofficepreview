package services

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
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

func TestLoadDiscoveryXml(t *testing.T) {
	fileData, _ := ioutil.ReadFile(exampleDiscoveryXmlPath)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(fileData)
	}))
	defer ts.Close()

	loadedData, err := LoadDiscoveryXml(ts.URL)

	assert := assert.New(t)
	assert.Equal(fileData, loadedData)
	assert.Nil(err)
}

func TestWopidDiscoveryGetXML(t *testing.T) {
	fileData, _ := ioutil.ReadFile(exampleDiscoveryXmlPath)

	wopiDiscovery1, err1 := ParseDiscoveryXml(fileData)
	wopiDiscovery2, err2 := ParseDiscoveryXml(wopiDiscovery1.GetXML())

	assert := assert.New(t)

	assert.Nil(err1)
	assert.Nil(err2)
	assert.NotNil(wopiDiscovery1)
	assert.NotNil(wopiDiscovery2)
	assert.Equal(wopiDiscovery1.GetXML(), wopiDiscovery2.GetXML())
}

func TestParseDiscoveryXmlUrl(t *testing.T) {
	fileData, _ := ioutil.ReadFile(exampleDiscoveryXmlPath)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(fileData)
	}))
	defer ts.Close()

	wopiDiscovery1, err1 := ParseDiscoveryXml(fileData)
	wopiDiscovery2, err2 := ParseDiscoveryXmlUrl(ts.URL)

	assert := assert.New(t)
	assert.Nil(err1)
	assert.Nil(err2)
	assert.NotNil(wopiDiscovery1)
	assert.NotNil(wopiDiscovery2)
	assert.Equal(wopiDiscovery1.GetXML(), wopiDiscovery2.GetXML())
}

func TestFindPreviewUrl(t *testing.T) {
	fileData, _ := ioutil.ReadFile(exampleDiscoveryXmlPath)
	wopiDiscovery, err1 := ParseDiscoveryXml(fileData)
	urlsrc, err2 := wopiDiscovery.FindPreviewUrl("internal-https", "xlsx")

	assert := assert.New(t)
	assert.Nil(err1)
	assert.NotNil(wopiDiscovery)
	assert.Nil(err2)
	assert.NotNil(urlsrc)
}
