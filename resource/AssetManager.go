package resource

import (
	"github.com/astaxie/beego/logs"
)

type AssetManager struct {
	*DataManager
}

func NewAssetManager(id string) *AssetManager {
	return &AssetManager{DataManager: &DataManager{ID: id, objAsset: make(map[string]interface{})}}

}

func (this *AssetManager) AddAsset(key string, asset interface{}) bool {

	if this.objAsset[key] == nil {
		this.objAsset[key] = &asset
		return true
	}
	logs.Error("AssetManager : AddAsset : key exist : ", key)
	return false
}

func (this *AssetManager) GetAsset(key string) interface{} {
	return this.objAsset[key]
}

func (this *AssetManager) HasAsset(key string) bool {
	return this.objAsset[key] != nil
}

func (this *AssetManager) RemoveAsset(key string) bool {

	if this.HasAsset(key) {
		delete(this.objAsset, key)
		return true
	}

	return false
}
