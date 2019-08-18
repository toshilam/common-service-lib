package resource

type DataManager struct {
	ID string

	//
	objAsset map[string]interface{}
}

func NewDataManager(id string) *DataManager {
	d := &DataManager{ID: id}
	d.objAsset = make(map[string]interface{})

	return d
}

func (this *DataManager) AddAsset(key string, asset interface{}) bool {
	return false
}

func (this *DataManager) GetAsset(key string) interface{} {
	return nil
}

func (this *DataManager) HasAsset(key string) bool {
	return false
}

func (this *DataManager) RemoveAsset(key string) bool {
	return false
}
