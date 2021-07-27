package indexModel

import (
	"app/config"
)

type ImgsModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Tablename      string   `json:"tablename"` 
	Object_id      uint   `json:"object_id"` 
	Imgurl      string   `json:"imgurl"` 
	Title      string   `json:"title"` 
	Orderindex      uint   `json:"orderindex"` 
	Userid      uint   `json:"userid"` 

}

func (ImgsModel) TableName() string {
	return config.TablePre + "imgs"
}

func ImgsList(list []ImgsModel) []ImgsModel {
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