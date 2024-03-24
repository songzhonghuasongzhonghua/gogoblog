package ctype

import "encoding/json"

type Role int

const (
	PermissionAdmin       Role = 1 //管理员
	PermissionUser        Role = 2 //普通用户
	PermissionVisitor     Role = 3 //游客
	PermissionDisableUser Role = 4 //禁言用户
)

func (role Role) MarshalJson() ([]byte, error) {
	return json.Marshal(role.String())
}

func (role Role) String() string {
	var str string
	switch role {
	case PermissionAdmin:
		str = "管理员"
	case PermissionUser:
		str = "普通用户"
	case PermissionVisitor:
		str = "游客"
	case PermissionDisableUser:
		str = "禁言用户"
	default:
		str = "其他"
	}

	return str
}
