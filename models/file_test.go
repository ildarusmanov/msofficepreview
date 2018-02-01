package models

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestCreateFile(t *testing.T) {
    filePath := "filePath"

    file := CreateFile(filePath)

    assert := assert.New(t)
    assert.NotNil(file)
    assert.Equal(file.FilePath, filePath)
}