package indexModel

import (
	"app/config"
)

type UserRankModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Rank_name      string   `json:"rank_name"` 
	Min_grade      uint   `json:"min_grade"` 
	Max_grade      uint   `json:"max_grade"` 
	Discount      uint   `json:"discount"` 
	Logo      string   `json:"logo"` 

}

func (UserRankModel) TableName() string {
	return config.TablePre + "user_rank"
}

func UserRankList(list []UserRankModel) []UserRankModel {
	slen := len(list)
	if slen == 0 {
		return list
	}

	for i := 0; i < slen; i++ {
		m := list[i]
		m.Logo = config.Image_site(m.Logo)
		list[i] = m
	}
	return list
}