package indexModel

import (
	"app/config"
)

type FavModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Tablename      string   `json:"tablename"` 
	Objectid      uint   `json:"objectid"` 
	Userid      uint   `json:"userid"` 

}

func (FavModel) TableName() string {
	return config.TablePre + "fav"
}

func FavList(list []FavModel) []FavModel {
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