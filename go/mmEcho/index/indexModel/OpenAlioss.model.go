package indexModel

import (
	"app/config"
)

type OpenAliossModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Title      string   `json:"title"` 
	Apptype      string   `json:"apptype"` 
	Createtime      string   `json:"createtime"` 
	Appid      string   `json:"appid"` 
	Appkey      string   `json:"appkey"` 
	Endpoint      string   `json:"endpoint"` 
	Bucket      string   `json:"bucket"` 
	Status      uint   `json:"status"` 

}

func (OpenAliossModel) TableName() string {
	return config.TablePre + "open_alioss"
}

func OpenAliossList(list []OpenAliossModel) []OpenAliossModel {
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