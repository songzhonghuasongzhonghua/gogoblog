package ctype

import "encoding/json"

type SignStatus int

const (
	SignQQ    SignStatus = 1
	SignGitee SignStatus = 2
	SignEmail SignStatus = 3
)

func (sign SignStatus) MarshalJson() ([]byte, error) {
	return json.Marshal(sign.String())
}

func (sign SignStatus) String() string {
	var str string
	switch sign {
	case SignQQ:
		str = "QQ"
	case SignGitee:
		str = "Gitee"
	case SignEmail:
		str = "Email"

	default:
		str = "其他"
	}

	return str
}
