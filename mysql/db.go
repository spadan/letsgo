package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	Db  *gorm.DB
	err error
)

func init() {
	Db, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
}

type MyUser struct {
	Id   uint
	Name string
	Age  uint8
}

type MyUser2 struct {
	Id   uint
	Name string
}

func Try() {
	/*Db.AutoMigrate(&MyUser{})
	u1 := MyUser{1, "xiongliang", 18}
	u2 := MyUser{2, "xiongxing", 17}
	Db.Create(&u1).Create(&u2)*/
	u3, u4 := MyUser2{}, MyUser2{}
	Db.Table("my_users").Select("id,name").First(&u3).Find(&u4, "id=?", 2)
	//Db.Table("my_users").Select([]string{"id", "name"}).Find(&u4, "id=?", 2)
	fmt.Println(u3, u4)
}
