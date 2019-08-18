package interfaces

type IVO interface {
	Clear() bool
	Clone() IVO
	ToJsonString() string
	GetType(structure interface{}) string
}
