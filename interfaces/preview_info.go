package interfaces

type PreviewInfo interface {
    GetSrc() string
    GetToken() string
    GetTokenTtl() int64
}
