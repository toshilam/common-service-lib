package data

type AuthVO struct {
	*VO
	Id       string //user id
	Key      string
	Name     string
	Platform string
	Secret   string
	Scope    string
	Tid      string
	Cid      string
}

func NewAuthVO(id string) *AuthVO {
	return &AuthVO{VO: &VO{ID: id}}
}

/* func (this *AuthVO) Clear() bool {
	return false
}

func (this *AuthVO) Clone() interfaces.IVO {
	return nil
} */
