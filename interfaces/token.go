package interfaces

type Token interface {
    GetValue() string
    GetTtl() int64
    GetFilePath() string
}
