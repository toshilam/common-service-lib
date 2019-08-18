package interfaces

type IDataManager interface {
	AddAsset(key string, asset interface{}) bool
	GetAsset(key string) interface{}
	HasAsset(key string) bool
	RemoveAsset(key string) bool
}
