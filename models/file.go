package models

type File struct {
	FilePath string `json:"file_path" binding:"required"`
}

func CreateFile(filePath string) *File {
	return &File{filePath}
}
