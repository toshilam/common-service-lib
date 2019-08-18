package interfaces

import "github.com/graphql-go/graphql"

type IGQLType interface {
	ToGQLObject() *graphql.Object
}
