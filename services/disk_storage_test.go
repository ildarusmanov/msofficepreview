package services

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "runtime"
  "path"
  "os"
)

var testFileId = "test_file.txt"

// helper function to build disk root path
// test root path should be relative to 
// current app directory
func getDiskRootPath() string {
  _, filename, _, _ := runtime.Caller(0)
  
  return path.Join(path.Dir(filename), "../test/tmp")
}

func TestCreateDiskStorage(t *testing.T) {
    storage := CreateDiskStorage(getDiskRootPath())

    assert.NotNil(t, storage)
}

func TestGetFileInfo(t *testing.T) {
    var (
        diskRootPath = getDiskRootPath()
        testFilePath = path.Join(diskRootPath, testFileId)
        testFileBaseName = path.Base(testFilePath)
        assert = assert.New(t)
        storage = CreateDiskStorage(diskRootPath)
    )

    f, err := os.Open(testFilePath)
    assert.Nil(err)
    defer f.Close()
    fstat, err := f.Stat()
    assert.Nil(err)

    fileInfo, err := storage.GetFileInfo(testFileId)

    assert.Nil(err)
    assert.Equal(fileInfo.GetFileName(), testFileBaseName)
    assert.Equal(fileInfo.GetSize(), fstat.Size())
    assert.Equal(fileInfo.GetVersion(), string(fstat.ModTime().Unix()))
}

func TestGetContents(t *testing.T) {
    storage := CreateDiskStorage(getDiskRootPath())

    fileContents, err := storage.GetContents(testFileId)

    assert := assert.New(t)
    assert.Nil(err)
    assert.NotNil(fileContents)
}
