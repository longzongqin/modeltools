package models

type UserInfo struct {
	Id int `json:"id"` 
	Nick string `json:"nick"` // 昵称
	HeadImg string `json:"head_img"` // 头像
	MiniOpenId string `json:"mini_open_id"` // 小程序openid
	Sex int8 `json:"sex"` // 0未知，1男，2女
	CrateTime time.Time `json:"crate_time"` 
	Status int8 `json:"status"` // 状态：0正常
	UserName string `json:"user_name"` // 用户名
	Password string `json:"password"` // 密码
}