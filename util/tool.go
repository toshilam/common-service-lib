package util

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/lithammer/shortuuid"

	"golang.org/x/crypto/bcrypt"
)

//base64
func Base64(text string, encode bool) string {
	if encode {
		return base64.StdEncoding.EncodeToString([]byte(text))
	}

	result, _ := base64.URLEncoding.DecodeString(text)
	return string(result)
}

//MD5 : hashing
func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

//HashPassword :
func HashPassword(text string) (string, error) {
	h, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		logs.Error("utils : HashPassword : error hashing password : ", err)
		return "", err
	}

	return string(h), err
}

func GenerateToken() string {
	return GenerateSecert("") //strings.ToUpper(MD5(shortuuid.New()))
}

func GenerateSecert(key string) string {
	return MD5(shortuuid.New() + key)
}

func GetUTCTime() time.Time {
	t := time.Now()
	localLocation, err := time.LoadLocation("UTC")
	if err != nil {
		fmt.Println(err)
	}
	timeUTC := t.In(localLocation)
	return timeUTC
}

func ToJsonString(o interface{}) string {
	result, err := json.Marshal(o)
	if err != nil {
		return ""
	}
	return string(result)
}

func GetType(structure interface{}) string {

	if t := reflect.TypeOf(structure); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}

func IsPointer(data interface{}) bool {
	return reflect.TypeOf(data).Kind() == reflect.Ptr
}

func GetKind(data interface{}) reflect.Kind {
	if IsPointer(data) {
		return reflect.TypeOf(data).Elem().Kind()
	}
	return reflect.TypeOf(data).Kind()
}

func IsStruct(data interface{}) bool {
	return GetKind(data) == reflect.Struct
}

func IsString(data interface{}) bool {
	return GetKind(data) == reflect.String
}

func IsInt(data interface{}) bool {
	return GetKind(data) == reflect.Int
}

func IsBool(data interface{}) bool {
	return GetKind(data) == reflect.Bool
}

func IsMap(data interface{}) bool {
	return GetKind(data) == reflect.Map
}

func encodeURIComponent(data interface{}) interface{} {

	if IsString(data) {
		return QueryEscape(data.(string))
	} /* else if IsMap(data) {
		for k, v := range data.(map[interface{}]interface{}) {

			if IsString(v) || IsMap(v) {

				data[k] = encodeURIComponent(v)
			}
		}
	} */

	return data
}

func QueryEscape(data string) string {
	r := url.QueryEscape(data)
	r = strings.Replace(r, "+", "%20", -1)
	return r
}

func QueryUnescape(data string) string {
	r, err := url.QueryUnescape(data)
	if err != nil {
		logs.Error("Tool : QueryUnescape : fail ", err)
		return data
	}

	return r
}
