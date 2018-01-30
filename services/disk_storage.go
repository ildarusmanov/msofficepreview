package services

type diskFileInfo struct {
    name string
    size int64
    version string
    ownerId string
}

func createDiskFileInfo(name string, size int64, version string, ownerId string) *diskFileInfo {
    return &diskFileInfo{
        name: name,
        size: size,
        version: version,
        ownerId: ownerId,
    }
}

func (f *diskFileInfo) GetFileName() string {
    return f.name
}

func (f *diskFileInfo) GetSize() int64 {
    return f.size
}

func (f *diskFileInfo) GetVersion() string {
    return f.version
}

func (f *diskFileInfo) GetOwnerId() string {
    return f.ownerId
}

type DiskStorage struct {
    rootPath string
}

func CreateDiskStorage(rootPath string) *DiskStorage {
    return &DiskStorage{rootPath: rootPath}
}

func (s *DiskStorage) GetFileInfo(fileId string) (*diskFileInfo, error) {
    return &diskFileInfo{}, nil
}

func (s *DiskStorage) GetContents(fileId string) ([]byte, error) {
    return []byte("test"), nil
}
