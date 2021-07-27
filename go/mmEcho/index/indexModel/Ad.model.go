package indexModel

import (
	"app/config"
)

type AdModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Tag_id      uint   `json:"tag_id"` 
	Title      string   `json:"title"` 
	Info      string   `json:"info"` 
	Link1      string   `json:"link1"` 
	Link2      string   `json:"link2"` 
	Starttime      uint   `json:"starttime"` 
	Endtime      uint   `json:"endtime"` 
	Imgurl      string   `json:"imgurl"` 
	Imgurl2      string   `json:"imgurl2"` 
	Orderindex      uint   `json:"orderindex"` 
	Status      uint   `json:"status"` 
	Dateline      uint   `json:"dateline"` 
	Price      float64   `json:"price"` 
	Object_id      uint   `json:"object_id"` 
	Tag_id_2nd      uint   `json:"tag_id_2nd"` 
	Checkbox_attr      string   `json:"checkbox_attr"` 

}

func (AdModel) TableName() string {
	return config.TablePre + "ad"
}

func AdList(list []AdModel) []AdModel {
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