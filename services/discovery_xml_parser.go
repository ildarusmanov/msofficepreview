package services

import (
    "encoding/xml"
)

type Action struct {
    Name string `xml:"name,attr"`
    Ext string `xml:"ext,attr"`
    Urlsrc string `xml:"urlsrc,attr"`
}

type App struct {
    Name string `xml:"name,attr"`
    Icon string `xml:"favIconUrl,attr"`
    License string `xml:"checkLicense,attr"`
    Actions []Action `xml:"action"`
}

type NetZone struct {
    Name string `xml:"name,attr"`
    Apps []App `xml:"app"`
}

type ProofKey struct {
    Oldvalue string `xml:"oldvalue,attr"`
    Oldmodulus string `xml:"oldmodulus,attr"`
    Oldexponent string `xml:"oldexponent,attr"`
    Value string `xml:"value,attr"`
    Modulus string `xml:"modulus,attr"`
    Exponent string `xml:"exponent,attr"`
}

type WopiDiscovery struct {
    XMLName xml.Name `xml:"wopi-discovery"`
    NetZones []NetZone `xml:"net-zone"`
    Proof ProofKey `xml:"proof-key"`
}

func CreateWopiDiscovery() *WopiDiscovery {
  return &WopiDiscovery{}
}

func ParseDiscoveryXml(data []byte) (*WopiDiscovery, error) {
    wopiDiscovery := CreateWopiDiscovery()

    if err := xml.Unmarshal(data, wopiDiscovery); err != nil {
        return nil, err
    }

    return wopiDiscovery, nil
}
