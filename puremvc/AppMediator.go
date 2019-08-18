package puremvc

import (
	"github.com/todo/common-service-lib/interfaces"
	"github.com/puremvc/puremvc-go-multicore-framework/src/patterns/mediator"
)

type AppMediator struct {
	*mediator.Mediator
	AppMain interfaces.IModuleMain
}

func (this *AppMediator) GetAssetManager() interfaces.IDataManager {
	return this.AppMain.GetAssetManager()
}

func (this *AppMediator) GetVOManager() interfaces.IDataManager {
	return this.AppMain.GetVOManager()
}
