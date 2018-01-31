package services

import (
    "encoding/xml"
    "io/ioutil"
    "net/http"
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

func (d *WopiDiscovery) GetXML() []byte {
    xmlData, err := xml.Marshal(d)

    if err != nil {
        return []byte{}
    }

    return xmlData
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
