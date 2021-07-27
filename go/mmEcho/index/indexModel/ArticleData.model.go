package indexModel

import (
	"app/config"
)

type ArticleDataModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Content      string   `json:"content"` 
	Createtime      string   `json:"createtime"` 

}

func (ArticleDataModel) TableName() string {
	return config.TablePre + "article_data"
}

func ArticleDataList(list []ArticleDataModel) []ArticleDataModel {
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