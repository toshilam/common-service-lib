package interfaces

import "github.com/puremvc/puremvc-go-multicore-framework/src/interfaces"

type IModuleMain interface {
	Setup(vo interface{}) bool
	GetAssetManager() IDataManager
	GetVOManager() IDataManager
	GetFacade() interfaces.IFacade
}
