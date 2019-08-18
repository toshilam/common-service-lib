package data

import (
	"github.com/todo/common-service-lib/interfaces"
	"github.com/todo/common-service-lib/util"
)

type VO struct {
	ID string `json:"id"`
}

func NewVO(id string) VO {
	return VO{ID: id}
}

func (this *VO) Clear() bool {
	return false
}

func (this *VO) Clone() interfaces.IVO {
	return nil
}

func (this *VO) ToJsonString() string {
	return util.ToJsonString(this)
}

/* func (this *VO) toJsonString(o interfaces.IVO) string {
	logs.Info("VO : toJsonString : ", o)
	result, err := json.Marshal(o)
	if err != nil {
		return ""
	}
	return string(result)
} */

func (this *VO) GetType(structure interface{}) string {

	return util.GetType(structure)
}

/* func (this *VO) GetSelf(structure interface{}) string {

    v := reflect.ValueOf(structure)
	numItem := v.NumField()
    values := make([]interface{}, numItem)

    for i := 0; i < numItem; i++ {
		v.Field(i).Type()
        values[i] = v.Field(i).Interface()
    }

     fmt.Println(values)
}
*/
