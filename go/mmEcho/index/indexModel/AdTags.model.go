package indexModel

import (
	"app/config"
)

type AdTagsModel struct {
	Tag_id      uint   `gorm:"primaryKey";json:"tag_id"`
	Title      string   `json:"title"` 
	Orderindex      uint   `json:"orderindex"` 
	Pid      uint   `json:"pid"` 
	Status      uint   `json:"status"` 
	Dateline      uint   `json:"dateline"` 
	M      string   `json:"m"` 
	A      string   `json:"a"` 
	Width      uint   `json:"width"` 
	Height      uint   `json:"height"` 
	Tagno      string   `json:"tagno"` 

}

func (AdTagsModel) TableName() string {
	return config.TablePre + "ad_tags"
}

func AdTagsList(list []AdTagsModel) []AdTagsModel {
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