package indexModel

import (
	"app/config"
)

type NavbarModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Title      string   `json:"title"` 
	Orderindex      uint   `json:"orderindex"` 
	Link_url      string   `json:"link_url"` 
	Target      string   `json:"target"` 
	Pid      uint   `json:"pid"` 
	Group_id      uint   `json:"group_id"` 
	M      string   `json:"m"` 
	A      string   `json:"a"` 
	Status      uint   `json:"status"` 
	Logo      string   `json:"logo"` 
	Icon      string   `json:"icon"` 
}

type NavbarTree struct {
	Id      uint   `json:"id"`
	Title      string   `json:"title"` 
	Orderindex      uint   `json:"orderindex"` 
	Link_url      string   `json:"link_url"` 
	Target      string   `json:"target"` 
	Pid      uint   `json:"pid"` 
	Group_id      uint   `json:"group_id"` 
	M      string   `json:"m"` 
	A      string   `json:"a"` 
	Status      uint   `json:"status"` 
	Logo      string   `json:"logo"` 
	Icon      string   `json:"icon"` 
	Child []NavbarTree	`json:"child"` 
}

func (NavbarModel) TableName() string {
	return config.TablePre + "navbar"
}

func NavbarList(list []NavbarModel) []NavbarModel {
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

func NavbarChild(list []NavbarModel) []NavbarTree {
	pList := []NavbarTree{}
	ilen := len(list)
	for i := 0; i < ilen; i++ {
		if list[i].Pid == 0 {
			item:=NavbarTree{
				Id:list[i].Id,
				Title:list[i].Title,
				Orderindex:list[i].Orderindex , 
				Link_url:list[i].Link_url , 
				Target:list[i].Target, 
				Pid :list[i].Pid, 
				Group_id :list[i].Group_id, 
				M :list[i].M, 
				A:list[i].A, 
				Status:list[i].Status, 
				Logo:list[i].Logo, 
				Icon:list[i].Icon,  
				Child: []NavbarTree{}, 
			}
			pList = append(pList, item)
		}
	}
	plen := len(pList)
	for j := 0; j < plen; j++ {
		child := []NavbarTree{}
		for i := 0; i < ilen; i++ {
			if pList[j].Id == list[i].Pid {
				item:=NavbarTree{
					Id:list[i].Id,
					Title:list[i].Title,
					Orderindex:list[i].Orderindex , 
					Link_url:list[i].Link_url , 
					Target:list[i].Target, 
					Pid :list[i].Pid, 
					Group_id :list[i].Group_id, 
					M :list[i].M, 
					A:list[i].A, 
					Status:list[i].Status, 
					Logo:list[i].Logo, 
					Icon:list[i].Icon,  
					Child: []NavbarTree{}, 
				}
				child = append(child, item)
			}
		}
		pList[j].Child = child

	}
	return pList

}