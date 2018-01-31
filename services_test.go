package main

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TODO add tests
func TestBuildServiceLocator(t *testing.T) {
	fileData, _ := ioutil.ReadFile("./test/fixtures/discovery.xml")

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(fileData)
	}))
	defer ts.Close()

	config := CreateNewConfig()
	config.DiscoveryXmlUrl = ts.URL

	locator, err := BuildServiceLocator(config)

	assert.Nil(t, err)
	assert.NotNil(t, locator)
}
