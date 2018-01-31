package interfaces

type WopiDiscovery interface {
    FindPreviewUrl(zone, ext string) (string, error)
}
