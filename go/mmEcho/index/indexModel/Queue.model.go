package indexModel

import (
	"app/config"
)

type QueueModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Orderindex      uint   `json:"orderindex"` 
	K      string   `json:"k"` 
	Content      string   `json:"content"` 

}

func (QueueModel) TableName() string {
	return config.TablePre + "queue"
}

func QueueList(list []QueueModel) []QueueModel {
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