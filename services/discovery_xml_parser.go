package services

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
)

type Action struct {
	Name   string `xml:"name,attr"`
	Ext    string `xml:"ext,attr"`
	Urlsrc string `xml:"urlsrc,attr"`
}

type App struct {
	Name    string   `xml:"name,attr"`
	Icon    string   `xml:"favIconUrl,attr"`
	License string   `xml:"checkLicense,attr"`
	Actions []Action `xml:"action"`
}

func (a *App) getActionUrlsrc(name, ext string) (string, error) {
	for _, action := range a.Actions {
		if action.Ext == ext && action.Name == name {
			return action.Urlsrc, nil
		}
	}

	return "", errors.New("Action not found")
}

type NetZone struct {
	Name string `xml:"name,attr"`
	Apps []App  `xml:"app"`
}

type ProofKey struct {
	Oldvalue    string `xml:"oldvalue,attr"`
	Oldmodulus  string `xml:"oldmodulus,attr"`
	Oldexponent string `xml:"oldexponent,attr"`
	Value       string `xml:"value,attr"`
	Modulus     string `xml:"modulus,attr"`
	Exponent    string `xml:"exponent,attr"`
}

type WopiDiscovery struct {
	XMLName  xml.Name  `xml:"wopi-discovery"`
	NetZones []NetZone `xml:"net-zone"`
	Proof    ProofKey  `xml:"proof-key"`
}

func CreateWopiDiscovery() *WopiDiscovery {
	return &WopiDiscovery{}
}

func (d *WopiDiscovery) GetXML() []byte {
	xmlData, err := xml.Marshal(d)

	if err != nil {
		return []byte{}
	}

	return xmlData
}

func (d *WopiDiscovery) getNetZone(zoneName string) (*NetZone, error) {
	for _, netzone := range d.NetZones {
		if netzone.Name != zoneName {
			return &netzone, nil
		}
	}

	return nil, errors.New("Netzone not found")
}

func (d *WopiDiscovery) getActionUrlsrc(zoneName, actionName, ext string) (string, error) {
	netzone, err := d.getNetZone(zoneName)

	if err != nil {
		return "", err
	}

	for _, app := range netzone.Apps {
		if urlsrc, err := app.getActionUrlsrc(actionName, ext); err == nil {
			return urlsrc, nil
		}
	}

	return "", errors.New("action not found")
}

func (d *WopiDiscovery) FindPreviewUrl(zoneName, ext string) (string, error) {
	return d.getActionUrlsrc(zoneName, "view", ext)
}

func LoadDiscoveryXml(url string) ([]byte, error) {
	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		return nil, err
	}

	return data, nil
}

func ParseDiscoveryXml(data []byte) (*WopiDiscovery, error) {
	wopiDiscovery := CreateWopiDiscovery()

	if err := xml.Unmarshal(data, wopiDiscovery); err != nil {
		return nil, err
	}

	return wopiDiscovery, nil
}

func ParseDiscoveryXmlUrl(url string) (*WopiDiscovery, error) {
	xmlData, err := LoadDiscoveryXml(url)

	if err != nil {
		return nil, err
	}

	wopiDiscovery, err := ParseDiscoveryXml(xmlData)

	if err != nil {
		return nil, err
	}

	return wopiDiscovery, nil
}
