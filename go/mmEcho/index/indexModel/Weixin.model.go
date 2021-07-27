package indexModel

import (
	"app/config"
)

type WeixinModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Token      string   `json:"token"` 
	Title      string   `json:"title"` 
	Dateline      uint   `json:"dateline"` 
	Status      uint   `json:"status"` 
	Userid      uint   `json:"userid"` 
	Domain      string   `json:"domain"` 
	Catid      uint   `json:"catid"` 
	Logo      string   `json:"logo"` 
	Imgurl      string   `json:"imgurl"` 
	Isrecommend      uint   `json:"isrecommend"` 
	Imgsdata      string   `json:"imgsdata"` 
	Appid      string   `json:"appid"` 
	Appkey      string   `json:"appkey"` 
	Isshow      uint   `json:"isshow"` 
	Ysid      string   `json:"ysid"` 
	Shopid      uint   `json:"shopid"` 
	Wx_username      string   `json:"wx_username"` 
	Wx_pwd      string   `json:"wx_pwd"` 
	Enkey      string   `json:"enkey"` 
	Mchid      string   `json:"mchid"` 
	Mchkey      string   `json:"mchkey"` 
	Sslcert_path      string   `json:"sslcert_path"` 
	Sslkey_path      string   `json:"sslkey_path"` 
	Openlogin      uint   `json:"openlogin"` 

}

func (WeixinModel) TableName() string {
	return config.TablePre + "weixin"
}

func WeixinList(list []WeixinModel) []WeixinModel {
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