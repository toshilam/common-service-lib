package data

type HttpVO struct {
	*VO
	BaseUrl string
	Timeout int // in Millisecond
	Headers map[string]string
}

func NewHttpVO(id, baseUrl string, timeout int /* , headers *map[string]string */) *HttpVO {
	o := &HttpVO{VO: &VO{ID: id}}
	o.Timeout = timeout
	// o.Headers = headers
	return o
}

/* func (this *HttpVO) Clear() bool {
	return false
}

func (this *HttpVO) Clone() interfaces.IVO {
	return nil
} */
