package indexModel

import (
	"app/config"
)

type ApppushPlanModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Pid      uint   `json:"pid"` 
	Typeid      uint   `json:"typeid"` 
	Status      uint   `json:"status"` 
	Content      string   `json:"content"` 
	Dateline      uint   `json:"dateline"` 
	Issuccess      uint   `json:"issuccess"` 
	Errmsg      string   `json:"errmsg"` 

}

func (ApppushPlanModel) TableName() string {
	return config.TablePre + "apppush_plan"
}

func ApppushPlanList(list []ApppushPlanModel) []ApppushPlanModel {
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