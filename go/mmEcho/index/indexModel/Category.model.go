package indexModel

import (
	"app/config"
)

type CategoryModel struct {
	Catid      uint   `gorm:"primaryKey";json:"catid"`
	Tablename      string   `json:"tablename"` 
	Pid      uint   `json:"pid"` 
	Cname      string   `json:"cname"` 
	Orderindex      uint   `json:"orderindex"` 
	Type_id      uint   `json:"type_id"` 
	Cat_tpl      string   `json:"cat_tpl"` 
	List_tpl      string   `json:"list_tpl"` 
	Show_tpl      string   `json:"show_tpl"` 
	Title      string   `json:"title"` 
	Keywords      string   `json:"keywords"` 
	Description      string   `json:"description"` 
	Status      uint   `json:"status"` 
	Level      uint   `json:"level"` 
	Topic_num      uint   `json:"topic_num"` 
	Comment_num      uint   `json:"comment_num"` 
	Last_post      string   `json:"last_post"` 
	Logo      string   `json:"logo"` 

}

func (CategoryModel) TableName() string {
	return config.TablePre + "category"
}

func CategoryList(list []CategoryModel) []CategoryModel {
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