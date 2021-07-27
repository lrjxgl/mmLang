package indexModel

import (
	"app/config"
)

type WeixinMenuModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Userid      uint   `json:"userid"` 
	Wid      uint   `json:"wid"` 
	Pid      uint   `json:"pid"` 
	Orderindex      uint   `json:"orderindex"` 
	Title      string   `json:"title"` 
	W_type      string   `json:"w_type"` 
	W_key      string   `json:"w_key"` 
	W_url      string   `json:"w_url"` 
	Sc_id      uint   `json:"sc_id"` 
	Shopid      uint   `json:"shopid"` 
	Content      string   `json:"content"` 

}

func (WeixinMenuModel) TableName() string {
	return config.TablePre + "weixin_menu"
}

func WeixinMenuList(list []WeixinMenuModel) []WeixinMenuModel {
	slen := len(list)
	if slen == 0 {
		return list
	}

	for i := 0; i < slen; i++ {
		m := list[i]
		
		list[i] = m
	}
	return list
}