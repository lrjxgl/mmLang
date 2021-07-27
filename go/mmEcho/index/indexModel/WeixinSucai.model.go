package indexModel

import (
	"app/config"
)

type WeixinSucaiModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Title      string   `json:"title"` 
	Dateline      uint   `json:"dateline"` 
	Shopid      uint   `json:"shopid"` 
	Content      string   `json:"content"` 
	Status      uint   `json:"status"` 
	Pid      uint   `json:"pid"` 
	Imgurl      string   `json:"imgurl"` 
	Wid      uint   `json:"wid"` 
	Userid      uint   `json:"userid"` 
	Description      string   `json:"description"` 
	Linkurl      string   `json:"linkurl"` 

}

func (WeixinSucaiModel) TableName() string {
	return config.TablePre + "weixin_sucai"
}

func WeixinSucaiList(list []WeixinSucaiModel) []WeixinSucaiModel {
	slen := len(list)
	if slen == 0 {
		return list
	}

	for i := 0; i < slen; i++ {
		m := list[i]
		m.Imgurl = config.Image_site(m.Imgurl)
		list[i] = m
	}
	return list
}