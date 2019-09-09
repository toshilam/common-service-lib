package module

import (
	"testing"

	"github.com/astaxie/beego/orm"
	"github.com/aws/aws-lambda-go/events"
	"github.com/todo/common-service-lib/data"

	_ "github.com/go-sql-driver/mysql"
)

func TestBaseHandler(t *testing.T) {

	t.Run("Test createShop : expecting shop creation success", func(t *testing.T) {

		orm.RegisterDriver("mysql", orm.DRMySQL)
		orm.RegisterDataBase("default", "mysql", "")

		id := "createShopHandler"
		values := map[string]interface{}{
			"Id": "8",
		}
		requestVO := data.NewRequestVO(id, &events.APIGatewayProxyRequest{RequestContext: events.APIGatewayProxyRequestContext{Authorizer: values}})

		baseHandler := NewBaseHandler(requestVO)

		if baseHandler != nil {
			t.Log("createShopHandler_test : NewBaseHandler : done ")
		} else {
			t.Error("createShopHandler_test : NewBaseHandler : Error :", baseHandler)

		}

	})

}
