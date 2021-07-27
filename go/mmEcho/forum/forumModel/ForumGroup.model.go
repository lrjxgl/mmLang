package forumModel

import (
	"app/config"
)

type ForumGroupModel struct {
	Gid      uint   `gorm:"primaryKey";json:"gid"`
	Title      string   `json:"title"` 
	Imgurl      string   `json:"imgurl"` 
	Description      string   `json:"description"` 
	Orderindex      uint   `json:"orderindex"` 
	Status      uint   `json:"status"` 
	Topic_num      uint   `json:"topic_num"` 
	Comment_num      uint   `json:"comment_num"` 
	View_num      uint   `json:"view_num"` 
	Isrecommend      uint   `json:"isrecommend"` 

}

func (ForumGroupModel) TableName() string {
	return config.TablePre + "mod_forum_group"
}

func ForumGroupList(list []ForumGroupModel) []ForumGroupModel {
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