package module

import (
	"github.com/todo/common-service-lib/data"
)

type BaseProvider struct {
	*data.VO
	RequestVO *data.RequestVO
	// userVO    *data.UserVO
}

func NewBaseProvider(context *data.RequestVO) *BaseProvider {
	o := &BaseProvider{VO: &data.VO{ID: ""}, RequestVO: context}
	return o
}

func Add(data interface{}) (interface{}, error) {
	return nil, nil
}

func Get(key string) (interface{}, error) {
	return nil, nil
}

func Update(key string) bool {
	return false
}

func Delete(key string) bool {
	return false
}
