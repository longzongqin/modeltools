package models

type RoleInfo struct {
	Id int `json:"id"` 
	Name string `json:"name"` 
	AuthInfoIds string `json:"auth_info_ids"` 
	Status int8 `json:"status"` // 0正常，1禁用
}