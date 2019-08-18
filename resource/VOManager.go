package resource

import (
	"github.com/astaxie/beego/logs"
)

type VOManager struct {
	*DataManager
}

func NewVOManager(id string) *VOManager {
	v := &VOManager{DataManager: &DataManager{ID: id, objAsset: make(map[string]interface{})}}
	v.objAsset = make(map[string]interface{})
	return v

}

func (this *VOManager) AddAsset(key string, asset interface{}) bool {
	if this.objAsset[key] == nil {
		this.objAsset[key] = asset
		logs.Info("---VOManager : AddAsset : added : ", this.objAsset, key, asset)
		return true
	}
	logs.Error("VOManager : AddAsset : key exist : ", key)
	return false
}

func (this *VOManager) GetAsset(key string) interface{} {
	return this.objAsset[key]
}

func (this *VOManager) HasAsset(key string) bool {
	return this.objAsset[key] != nil
}

func (this *VOManager) RemoveAsset(key string) bool {

	if this.HasAsset(key) {
		delete(this.objAsset, key)
		return true
	}

	return false
}
