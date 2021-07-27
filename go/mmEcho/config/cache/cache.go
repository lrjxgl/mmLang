package cache

import (
	"app/config"
	"app/index/indexModel"
)

func CacheSet(k string, v string, expire int64) {
	var db = config.Db
	row := indexModel.DbcacheModel{}
	res := db.Where("k=?", k).First(&row)
	data := indexModel.DbcacheModel{
		K:      k,
		V:      v,
		Expire: uint(expire),
	}
	if res.Error != nil {

		db.Create(&data)
	} else {
		db.Where("k=?", k).Updates(&data)
	}

}
func CacheGet(k string) string {
	var db = config.Db
	row := indexModel.DbcacheModel{}
	res := db.Where("k=?", k).First(&row)
	if res.Error != nil {
		return ""
	}
	return row.V
}
