package indexModel

import (
	"app/config"
)

type ArticleModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Title      string   `json:"title"` 
	Catid      uint   `json:"catid"` 
	Catid_top      uint   `json:"catid_top"` 
	Catid_2nd      uint   `json:"catid_2nd"` 
	Love_num      uint   `json:"love_num"` 
	Fav_num      uint   `json:"fav_num"` 
	Description      string   `json:"description"` 
	Status      uint   `json:"status"` 
	Isrecommend      uint   `json:"isrecommend"` 
	Imgurl      string   `json:"imgurl"` 
	Createtime      string   `json:"createtime"` 
	Comment_num      uint   `json:"comment_num"` 
	Tpl      string   `json:"tpl"` 
	Author      string   `json:"author"` 
	Isnew      uint   `json:"isnew"` 
	Ishot      uint   `json:"ishot"` 
	View_num      uint   `json:"view_num"` 
	Price      float64   `json:"price"` 
	Market_price      float64   `json:"market_price"` 
	Total_num      uint   `json:"total_num"` 
	Sold_num      uint   `json:"sold_num"` 
	Grade      uint   `json:"grade"` 
	Downurl      string   `json:"downurl"` 
	Downsize      float64   `json:"downsize"` 
	Orderindex      uint   `json:"orderindex"` 
	Videourl      string   `json:"videourl"` 
	Isindex      uint   `json:"isindex"` 
	Tags      string   `json:"tags"` 
	Imgsdata      string   `json:"imgsdata"` 

}

func (ArticleModel) TableName() string {
	return config.TablePre + "article"
}

func ArticleList(list []ArticleModel) []ArticleModel {
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