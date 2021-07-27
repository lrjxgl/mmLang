package indexModel

import (
	"app/config"
)

type KefuMsgindexModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Createtime      string   `json:"createtime"` 
	Userid      uint   `json:"userid"` 
	Content      string   `json:"content"` 
	Status      uint   `json:"status"` 
	Isreply      uint   `json:"isreply"` 
	Isread      uint   `json:"isread"` 

}

func (KefuMsgindexModel) TableName() string {
	return config.TablePre + "kefu_msgindex"
}

func KefuMsgindexList(list []KefuMsgindexModel) []KefuMsgindexModel {
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