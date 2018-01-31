package services

import (
    "os"
    "path"
    "io/ioutil"
    "github.com/ildarusmanov/msofficepreview/interfaces"
)

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

func (s *DiskStorage) GetFileInfo(fileId string) (interfaces.FileInfo, error) {
    filePath := s.getFilePath(fileId)
    file, err := os.Open(filePath) 

    if err != nil {
        return nil, err
    }

    defer file.Close()

    fileInfo, err := file.Stat()

    if err != nil {
        return nil, err
    }

    return &diskFileInfo{
            name: path.Base(filePath),
            size: fileInfo.Size(),
            version: string(fileInfo.ModTime().Unix()),
            ownerId: "root",
        }, nil
}

func (s *DiskStorage) GetContents(fileId string) ([]byte, error) {
    filePath := s.getFilePath(fileId)
 
    return ioutil.ReadFile(filePath)
}

func (s *DiskStorage) getFilePath(fileId string) string {
    return path.Join(s.rootPath, fileId)
}
