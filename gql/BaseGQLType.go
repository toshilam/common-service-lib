package gql

import (
	"github.com/todo/common-service-lib/data"
	"github.com/astaxie/beego/logs"
	"github.com/graphql-go/graphql"
)

//BaseGQLType :
type BaseGQLType struct {
	*data.VO
	Name    string
	Context *data.RequestVO
	Fields  interface{}

	GQLObject *graphql.Object
}

//NewBaseGQLType :
func NewBaseGQLType(name string, context *data.RequestVO) *BaseGQLType {
	logs.Info("NewBaseGQLType : ", name)
	logs.Info(context)
	return &BaseGQLType{VO: &data.VO{ID: name}, Name: name, Context: context}
}

//GetFields : to be implemented by child struct
/* func (this *BaseGQLType) GetFields() map[string]*graphql.Field {
	return graphql.Fields{}
} */

//ToGQLObject :
func (this *BaseGQLType) ToGQLObject() *graphql.Object {

	logs.Info(this.Name, ": ToGQLObject : ", this.Fields)

	if this.GQLObject == nil {

		this.GQLObject = graphql.NewObject(graphql.ObjectConfig{
			Name:   this.Name,
			Fields: this.Fields,
		})
	}

	return this.GQLObject
}

/* func (this BaseGQLType) String() string {
	return fmt.Sprintf("[struct : %s]", this.Name)
} */
