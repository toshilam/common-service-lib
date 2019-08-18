package interfaces

type IProvider interface {
	Add(data interface{}) (interface{}, error)
	Get(key string) (interface{}, error)
	Update(key string) bool
	Delete(key string) bool
}
