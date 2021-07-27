package indexModel

import (
	"app/config"
)

type GradeLogModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Dateline      uint   `json:"dateline"` 
	Userid      uint   `json:"userid"` 
	Type_id      uint   `json:"type_id"` 
	Content      string   `json:"content"` 
	Isdelete      uint   `json:"isdelete"` 
	Grade      uint   `json:"grade"` 

}

func (GradeLogModel) TableName() string {
	return config.TablePre + "grade_log"
}

func GradeLogList(list []GradeLogModel) []GradeLogModel {
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