package interfaces

type ServiceLocator interface {
    Set(key string, value interface{})
    Get(key string) interface{}
}
