package indexModel

import (
	"app/config"
)

type WeixinCommandModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Wid      uint   `json:"wid"` 
	Title      string   `json:"title"` 
	Command      string   `json:"command"` 
	Type_id      uint   `json:"type_id"` 
	Content      string   `json:"content"` 
	Dateline      uint   `json:"dateline"` 
	Fun      string   `json:"fun"` 
	Isdefault      uint   `json:"isdefault"` 
	Userid      uint   `json:"userid"` 
	Shopid      uint   `json:"shopid"` 
	Sc_id      uint   `json:"sc_id"` 

}

func (WeixinCommandModel) TableName() string {
	return config.TablePre + "weixin_command"
}

func WeixinCommandList(list []WeixinCommandModel) []WeixinCommandModel {
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