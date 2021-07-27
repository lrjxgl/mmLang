package indexModel

import (
	"app/config"
)

type SeoModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	M      string   `json:"m"` 
	A      string   `json:"a"` 
	Cname      string   `json:"cname"` 
	Title      string   `json:"title"` 
	Description      string   `json:"description"` 
	Keywords      string   `json:"keywords"` 
	Object_id      uint   `json:"object_id"` 

}

func (SeoModel) TableName() string {
	return config.TablePre + "seo"
}

func SeoList(list []SeoModel) []SeoModel {
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