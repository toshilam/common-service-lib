package module

import (
	"errors"
	"reflect"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/validation"
	"github.com/graphql-go/graphql"
	"github.com/mitchellh/mapstructure"
	"github.com/todo/common-service-lib/data"
	"github.com/todo/common-service-lib/util"
)

type BaseHandler struct {
	*data.VO
	RequestVO *data.RequestVO
	AuthVO    *data.AuthVO
}

func NewBaseHandler(context *data.RequestVO) *BaseHandler {
	var authVO data.AuthVO
	decodeError := mapstructure.Decode(context.Request.RequestContext.Authorizer, &authVO)
	logs.Error(authVO)
	if decodeError != nil {
		logs.Error("BaseHandler : NewBaseHandler : fail decoding :", decodeError)
		panic(decodeError)
	}
	o := &BaseHandler{VO: &data.VO{ID: ""}, RequestVO: context, AuthVO: &authVO}
	return o
}

//Authrizer : to get current user data by passing data key
func (this *BaseHandler) Authorizer(key string) string {
	value, ok := this.RequestVO.Request.RequestContext.Authorizer[key].(string)
	if ok {
		return value
	}
	logs.Error("BaseHandler : Authorizer : unable to get ", key, ok, this.RequestVO.Request.RequestContext.Authorizer)
	return ""
}

func (this *BaseHandler) Validate(requestData interface{}, output interface{}) (interface{}, error) {

	err := mapstructure.Decode(requestData, output)
	if err == nil {

		logs.Info(this.GetType(this), ": Validate : Decoded struct :", this.GetType(output), output)

		//TODO : only support pointer struct? loop through array|map ?
		if util.IsPointer(output) {
			v := reflect.Indirect(reflect.ValueOf(output))
			for i := 0; i < v.NumField(); i++ {

				if util.IsString(v.Field(i).Interface()) {
					logs.Info(this.GetType(this), ": Validate : QueryUnescape : from ", v.Field(i).Interface(), ", to : ", util.QueryUnescape(v.Field(i).Interface().(string)))
					v.Field(i).SetString(util.QueryUnescape(v.Field(i).Interface().(string)))
				}

			}
		}

		valid := validation.Validation{}
		b, _ := valid.Valid(output)

		logs.Info(this.GetType(this), ": Validate : validation : success :", b)

		if b {
			return output, nil
		}

		msg := "Error :"
		for _, err := range valid.Errors {
			msg += err.Key + ", " + err.Message
		}
		err = errors.New(msg)

	}

	return nil, err
}

func (this *BaseHandler) Resolve(params graphql.ResolveParams) (interface{}, error) {

	return nil, nil
}

func (this *BaseHandler) ToGQLField() *graphql.Field {
	return nil
}

func (this *BaseHandler) Response(code string, result interface{}, err interface{}) *data.ResultVO {
	return data.NewResultVO(this.ID, code, result, err)
}
