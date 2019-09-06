package models

type AdminInfo struct {
	Id int `json:"id"` 
	Name string `json:"name"` 
	Username string `json:"username"` 
	Password string `json:"password"` 
	RoleInfoId int `json:"role_info_id"` // 角色
	Status int8 `json:"status"` // -1删除，0正常，1禁用
}