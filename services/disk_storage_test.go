package services

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "runtime"
  "path"
)

var testFileId = "test_file.txt"

// helper function to build disk root path
// test root path should be relative to 
// current app directory
func getDiskRootPath() string {
  _, filename, _, _ := runtime.Caller(0)
  
  return path.Join(path.Dir(filename), "test/tmp")
}

func TestCreateDiskStorage(t *testing.T) {
    storage := CreateDiskStorage(getDiskRootPath())

    assert.NotNil(t, storage)
}

func TestGetFileInfo(t *testing.T) {
    storage := CreateDiskStorage(getDiskRootPath())

    fileInfo := storage.GetFileInfo(testFileId)
    
    assert := assert.New(t)
    assert.Equal(fileInfo.GetFileName(), testFileBaseName)
    assert.Equal(fileInfo.GetSize(), testFileSize)
    assert.Equal(fileInfo.GetVersion(), testFileVersion)
    assert.Equal(fileInfo.GetOwnerId(), testFileOwnerId)
}

func TestGetContents(t *testing.T) {
    storage := CreateDiskStorage(getDiskRootPath())

    fileContents := storage.GetContents(testFileId)

    assert.NotNil(t, fileContents)
}
