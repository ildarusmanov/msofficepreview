package main

import (
	"github.com/ildarusmanov/msofficepreview/services"
)

func BuildServiceLocator(config *Config) (*services.ServiceLocator, error) {
	wopiDiscovery, err := services.ParseDiscoveryXmlUrl(config.DiscoveryXmlUrl)

	if err != nil {
		return nil, err
	}

	tokenProvider := services.CreateMemoryTokenProvider(config.TokenLifetime)
	storage := services.CreateDiskStorage(config.StorageRoot)
	previewGenerator := services.CreatePreviewGenerator(
		config.ServerHost,
		wopiDiscovery,
		tokenProvider,
		storage,
	)

	locator := services.CreateServiceLocator()
	locator.Set("TokenProvider", tokenProvider)
	locator.Set("Storage", storage)
	locator.Set("PreviewGenerator", previewGenerator)

	return locator, nil
}
