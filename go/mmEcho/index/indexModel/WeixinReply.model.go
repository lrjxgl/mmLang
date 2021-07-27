package indexModel

import (
	"app/config"
)

type WeixinReplyModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Status      uint   `json:"status"` 
	Openid      string   `json:"openid"` 
	Msgtype      string   `json:"msgtype"` 
	Createtime      uint   `json:"createtime"` 
	Content      string   `json:"content"` 
	Msgid      string   `json:"msgid"` 
	Picurl      string   `json:"picurl"` 
	Mediaid      string   `json:"mediaid"` 
	Thumbmediaid      string   `json:"thumbmediaid"` 
	Format      string   `json:"format"` 
	Location_x      string   `json:"location_x"` 
	Location_y      string   `json:"location_y"` 
	Scale      uint   `json:"scale"` 
	Label      string   `json:"label"` 
	Title      string   `json:"title"` 
	Description      string   `json:"description"` 
	Url      string   `json:"url"` 
	Wid      uint   `json:"wid"` 
	Shopid      uint   `json:"shopid"` 
	Fromusername      string   `json:"fromusername"` 
	Tousername      string   `json:"tousername"` 

}

func (WeixinReplyModel) TableName() string {
	return config.TablePre + "weixin_reply"
}

func WeixinReplyList(list []WeixinReplyModel) []WeixinReplyModel {
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