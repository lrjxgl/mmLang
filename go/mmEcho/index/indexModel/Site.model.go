package indexModel

import (
	"app/config"
)

type SiteModel struct {
	Siteid      uint   `gorm:"primaryKey";json:"siteid"`
	Sitename      string   `json:"sitename"` 
	Domain      string   `json:"domain"` 
	Is_open      uint   `json:"is_open"` 
	Title      string   `json:"title"` 
	Keywords      string   `json:"keywords"` 
	Description      string   `json:"description"` 
	Close_why      string   `json:"close_why"` 
	Logo      string   `json:"logo"` 
	Icp      string   `json:"icp"` 
	Status      uint   `json:"status"` 
	Template      string   `json:"template"` 
	Statjs      string   `json:"statjs"` 
	Index_template      string   `json:"index_template"` 
	Wap_template      string   `json:"wap_template"` 
	Wapbg      string   `json:"wapbg"` 

}

func (SiteModel) TableName() string {
	return config.TablePre + "site"
}

func SiteList(list []SiteModel) []SiteModel {
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