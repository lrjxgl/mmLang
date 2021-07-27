package indexModel

import (
	"app/config"
)

type KefuMsgModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Createtime      string   `json:"createtime"` 
	Userid      uint   `json:"userid"` 
	Content      string   `json:"content"` 
	Status      uint   `json:"status"` 
	Tablename      string   `json:"tablename"` 

}

func (KefuMsgModel) TableName() string {
	return config.TablePre + "kefu_msg"
}

func KefuMsgList(list []KefuMsgModel) []KefuMsgModel {
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