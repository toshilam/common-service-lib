package data

import (
	"github.com/todo/common-service-lib/interfaces"
	"github.com/todo/common-service-lib/resource"
	"github.com/astaxie/beego/orm"
	"github.com/aws/aws-lambda-go/events"
)

type RequestVO struct {
	*VO
	Request      *events.APIGatewayProxyRequest
	Channel      chan *ResultVO
	AssetManager *resource.AssetManager
	VOManager    *resource.VOManager
	ORM          orm.Ormer
}

func NewRequestVO(id string, request *events.APIGatewayProxyRequest) *RequestVO {
	o := &RequestVO{VO: &VO{ID: id}}

	o.Request = request
	o.Channel = make(chan *ResultVO)
	o.AssetManager = resource.NewAssetManager(id)
	o.VOManager = resource.NewVOManager(id)

	//create a orm object autometically
	o.ORM = orm.NewOrm()

	return o
}

func (this *RequestVO) Clone() interfaces.IVO {
	return NewRequestVO(this.ID, this.Request)
}
