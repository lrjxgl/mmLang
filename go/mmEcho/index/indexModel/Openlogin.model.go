package indexModel

import (
	"app/config"
)

type OpenloginModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Userid      uint   `json:"userid"` 
	Xfrom      string   `json:"xfrom"` 
	Openid      string   `json:"openid"` 
	Accesstoken      string   `json:"accesstoken"` 
	Createtime      string   `json:"createtime"` 
	Nickname      string   `json:"nickname"` 
	User_head      string   `json:"user_head"` 
	Gender      uint   `json:"gender"` 
	Unionid      string   `json:"unionid"` 

}

func (OpenloginModel) TableName() string {
	return config.TablePre + "openlogin"
}

func OpenloginList(list []OpenloginModel) []OpenloginModel {
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