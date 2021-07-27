package config;
import (
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
	 
)
var Db  *gorm.DB;
var TablePre="sky_";
func init(){
	fmt.Print("init mysql")
	dsn := "root:root@tcp(127.0.0.1:3306)/mmlang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	
	if err!= nil{
	    panic(err)
	}
	
	Db=db;
}
