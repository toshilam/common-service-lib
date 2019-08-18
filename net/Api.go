package net

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/todo/common-service-lib/data"
	"github.com/astaxie/beego/logs"
)

type Api struct {
	*data.VO
	HttpVO *data.HttpVO
}

func NewApi(config *data.HttpVO) *Api {
	logs.Info("Api : NewApi : ")

	return &Api{
		VO: &data.VO{ID: ""},
	}
}

func (this *Api) Post(path string, body []byte) (interface{}, error) {

	request, err := http.NewRequest(
		"POST",
		this.HttpVO.BaseUrl+path,
		bytes.NewBuffer(body),
	)

	if err != nil {
		logs.Error("Api : send : fail building request : ", err)
		return nil, err
	}

	if this.HttpVO.Headers != nil && len(this.HttpVO.Headers) > 0 {
		for k, v := range this.HttpVO.Headers {
			logs.Info("Api : Post : adding header : ", k, v)
			request.Header.Set(k, v)
		}
	}

	httpClient := &http.Client{Timeout: time.Duration(this.HttpVO.Timeout) * time.Millisecond}

	logs.Info("Api : send : requesting ", request.URL)
	response, err := httpClient.Do(request)
	if err != nil /* || response.StatusCode != 200 */ {

		// if e, ok := err.(net.Error); ok && e.Timeout() {
		// 	logs.Error("Api : send : timeout")
		// } else {
		logs.Error("Api : send : request failed: ", err)
		// }
		return nil, err
	}

	defer response.Body.Close()

	body, readErr := ioutil.ReadAll(response.Body)
	logs.Info("Api : send : reading response ", body, readErr)

	if readErr != nil {
		logs.Error("Api : send : response read failed: ", readErr)
		return nil, readErr
	}

	//TODO : return vo with response, body...
	return string(body), nil
}

func (this *Api) Get(path string) (interface{}, error) {

	httpClient := &http.Client{Timeout: time.Duration(this.HttpVO.Timeout) * time.Millisecond}
	req, err := httpClient.Get(this.HttpVO.BaseUrl + path)

	if err != nil {
		return nil, err
	}

	// if req.StatusCode != 200 {
	// 	return
	// }

	body, err := ioutil.ReadAll(req.Body)
	req.Body.Close()
	// defer rs.Body.Close()
	// log.LOG("body", string(body))

	//TODO : return vo with response, body...
	return string(body), nil
}
