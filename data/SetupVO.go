package data

import (
	"github.com/todo/common-service-lib/interfaces"
)

type SetupVO struct {
	*VO
	AssetManager interfaces.IDataManager
	VOManager    interfaces.IDataManager
}

func NewSetupVO(id string, assetManager interfaces.IDataManager, voManager interfaces.IDataManager) *SetupVO {
	return &SetupVO{VO: &VO{ID: id}, AssetManager: assetManager, VOManager: voManager}
}

func (this *SetupVO) Clear() bool {
	return false
}

func (this *SetupVO) Clone() interfaces.IVO {
	return nil
}
