package indexModel

import (
	"app/config"
)

type ArticleCommentModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Userid      uint   `json:"userid"` 
	Status      uint   `json:"status"` 
	Objectid      uint   `json:"objectid"` 
	Pid      uint   `json:"pid"` 
	Createtime      string   `json:"createtime"` 
	Content      string   `json:"content"` 
	Ip      string   `json:"ip"` 
	Ip_city      string   `json:"ip_city"` 
	Love_num      uint   `json:"love_num"` 
	Imgsdata      string   `json:"imgsdata"` 

}

func (ArticleCommentModel) TableName() string {
	return config.TablePre + "article_comment"
}

func ArticleCommentList(list []ArticleCommentModel) []ArticleCommentModel {
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