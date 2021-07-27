package indexModel

import (
	"app/config"
)

type WeixinUserModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Openid      string   `json:"openid"` 
	Dateline      uint   `json:"dateline"` 
	Add_time      uint   `json:"add_time"` 
	Last_time      uint   `json:"last_time"` 
	Del_time      uint   `json:"del_time"` 
	Status      uint   `json:"status"` 
	Nickname      string   `json:"nickname"` 
	Sex      uint   `json:"sex"` 
	City      string   `json:"city"` 
	Country      string   `json:"country"` 
	Province      string   `json:"province"` 
	User_head      string   `json:"user_head"` 
	Update_time      uint   `json:"update_time"` 
	Reply_num      uint   `json:"reply_num"` 
	Wid      uint   `json:"wid"` 
	Shopid      uint   `json:"shopid"` 

}

func (WeixinUserModel) TableName() string {
	return config.TablePre + "weixin_user"
}

func WeixinUserList(list []WeixinUserModel) []WeixinUserModel {
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