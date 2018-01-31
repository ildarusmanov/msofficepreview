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

    fileInfo1, err1 := storage.GetFileInfo(testFileId)
    fileInfo2, err2 := storage.GetFileInfo(testFileId + "1")

    assert.Nil(err1)
    assert.Equal(fileInfo1.GetFileName(), testFileBaseName)
    assert.Equal(fileInfo1.GetSize(), fstat.Size())
    assert.Equal(fileInfo1.GetVersion(), string(fstat.ModTime().Unix()))
    assert.NotNil(err2)
    assert.Nil(fileInfo2)
}

func TestGetContents(t *testing.T) {
    storage := CreateDiskStorage(getDiskRootPath())

    fileContent1, err1 := storage.GetContents(testFileId)
    fileContent2, err2 := storage.GetContents(testFileId + "1")

    assert := assert.New(t)
    assert.Nil(err1)
    assert.NotNil(fileContent1)
    assert.NotNil(err2)
    assert.Nil(fileContent2)
}
