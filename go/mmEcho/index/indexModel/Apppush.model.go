package indexModel

import (
	"app/config"
)

type ApppushModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Token      string   `json:"token"` 
	Clientid      string   `json:"clientid"` 
	Appname      string   `json:"appname"` 
	Appid      string   `json:"appid"` 
	Appkey      string   `json:"appkey"` 
	Dateline      uint   `json:"dateline"` 
	Userid      uint   `json:"userid"` 
	Shopadmin      uint   `json:"shopadmin"` 
	Openid      string   `json:"openid"` 

}

func (ApppushModel) TableName() string {
	return config.TablePre + "apppush"
}

func ApppushList(list []ApppushModel) []ApppushModel {
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