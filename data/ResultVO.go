package data

import (
	"github.com/todo/common-service-lib/interfaces"
	"github.com/todo/common-service-lib/util"
)

type ResultVO struct {
	*VO
	Code  string      `json:"code"`
	Data  interface{} `json:"data"`
	Error interface{} `json:"error"`
}

func NewResultVO(id string, code string, data interface{}, err interface{}) *ResultVO {
	o := &ResultVO{VO: &VO{ID: id}}
	o.Data = data
	o.Code = code
	o.Error = err
	return o
}

func (this *ResultVO) Clear() bool {
	this.Data = nil
	return true
}

func (this *ResultVO) Clone() interfaces.IVO {
	return NewResultVO(this.ID, this.Code, this.Data, this.Error)
}

func (this *ResultVO) ToJsonString() string {

	return util.ToJsonString(this)
}
