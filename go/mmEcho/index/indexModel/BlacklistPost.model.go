package indexModel

import (
	"app/config"
)

type BlacklistPostModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Userid      uint   `json:"userid"` 
	Siteid      uint   `json:"siteid"` 
	Dateline      uint   `json:"dateline"` 
	Etime      uint   `json:"etime"` 

}

func (BlacklistPostModel) TableName() string {
	return config.TablePre + "blacklist_post"
}

func BlacklistPostList(list []BlacklistPostModel) []BlacklistPostModel {
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