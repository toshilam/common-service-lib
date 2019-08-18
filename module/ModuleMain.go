package module

import (
	"github.com/todo/common-service-lib/data"
	"github.com/todo/common-service-lib/interfaces"
	"github.com/astaxie/beego/logs"
	i "github.com/puremvc/puremvc-go-multicore-framework/src/interfaces"
	"github.com/puremvc/puremvc-go-multicore-framework/src/patterns/facade"
)

type ModuleMain struct {
	Name    string
	SetupVO *data.SetupVO
	Facade  i.IFacade
}

func (this *ModuleMain) Setup(d interface{}) bool {

	vo := d.(*data.SetupVO)
	logs.Info("ModuleMain : Setup :", vo)

	if this.SetupVO == nil && vo.AssetManager != nil && vo.VOManager != nil {
		this.Facade = facade.GetInstance(vo.ID, func() i.IFacade { return &facade.Facade{Key: vo.ID} })
		this.SetupVO = vo
		return true
	}

	return false
}

func (this *ModuleMain) GetFacade() i.IFacade {
	return this.Facade
}

func (this *ModuleMain) GetAssetManager() interfaces.IDataManager {
	return this.SetupVO.AssetManager
}

func (this *ModuleMain) GetVOManager() interfaces.IDataManager {
	return this.SetupVO.VOManager
}
