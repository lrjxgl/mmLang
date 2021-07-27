package indexModel

import (
	"app/config"
)

type UserModel struct {
	Userid      uint   `gorm:"primaryKey";json:"userid"`
	Username      string   `json:"username"` 
	Telephone      string   `json:"telephone"` 
	Nickname      string   `json:"nickname"` 
	Money      float64   `json:"money"` 
	Gold      uint   `json:"gold"` 
	Grade      uint   `json:"grade"` 
	Createtime      string   `json:"createtime"` 
	Lasttime      string   `json:"lasttime"` 
	User_type      uint   `json:"user_type"` 
	Status      uint   `json:"status"` 
	Is_auth      uint   `json:"is_auth"` 
	User_head      string   `json:"user_head"` 
	Follow_num      uint   `json:"follow_num"` 
	Followed_num      uint   `json:"followed_num"` 
	Gender      uint   `json:"gender"` 
	Invite_userid      uint   `json:"invite_userid"` 
	Birthday      string   `json:"birthday"` 
	Description      string   `json:"description"` 

}

func (UserModel) TableName() string {
	return config.TablePre + "user"
}

func UserList(list []UserModel) []UserModel {
	slen := len(list)
	if slen == 0 {
		return list
	}

	for i := 0; i < slen; i++ {
		m := list[i]
		m.User_head = config.Image_site(m.User_head)
		list[i] = m
	}
	return list
}