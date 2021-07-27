package indexModel

import (
	"app/config"
)

type PagetplModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	M      string   `json:"m"` 
	A      string   `json:"a"` 
	Tpl      string   `json:"tpl"` 

}

func (PagetplModel) TableName() string {
	return config.TablePre + "pagetpl"
}

func PagetplList(list []PagetplModel) []PagetplModel {
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