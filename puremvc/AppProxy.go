package puremvc

import (
	"github.com/todo/common-service-lib/interfaces"
	"github.com/puremvc/puremvc-go-multicore-framework/src/patterns/proxy"
)

type AppProxy struct {
	*proxy.Proxy
	AppMain interfaces.IModuleMain
}

func (this *AppProxy) GetAssetManager() interfaces.IDataManager {
	return this.AppMain.GetAssetManager()
}

func (this *AppProxy) GetVOManager() interfaces.IDataManager {
	return this.AppMain.GetVOManager()
}
