package indexModel

import (
	"app/config"
)

type GuestModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Title      string   `json:"title"` 
	Type_id      uint   `json:"type_id"` 
	Userid      uint   `json:"userid"` 
	Nickname      string   `json:"nickname"` 
	Email      string   `json:"email"` 
	Telephone      string   `json:"telephone"` 
	Dateline      uint   `json:"dateline"` 
	Gender      uint   `json:"gender"` 
	Reply      string   `json:"reply"` 
	Isreply      uint   `json:"isreply"` 
	Reply_time      uint   `json:"reply_time"` 
	Money      float64   `json:"money"` 
	Address      string   `json:"address"` 
	Qq      string   `json:"qq"` 
	Num      uint   `json:"num"` 
	Status      uint   `json:"status"` 
	Content      string   `json:"content"` 

}

func (GuestModel) TableName() string {
	return config.TablePre + "guest"
}

func GuestList(list []GuestModel) []GuestModel {
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