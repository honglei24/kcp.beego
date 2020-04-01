package main

import (
	"github.com/astaxie/beego"
	_ "kcp/routers"
	_ "kcp/sysinit"
)

func main() {
	//o := orm.NewOrm()
	//
	//user := models.User{Id: 1, Name: "slene"}
	//
	//// insert
	//id, err := o.Insert(&user)
	//fmt.Printf("ID: %d, ERR: %v\n", id, err)
	//
	//// update
	//user.Name = "astaxie"
	//num, err := o.Update(&user)
	//fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	//// read one
	//u := models.User{Id: user.Id}
	//err = o.Read(&u)
	//fmt.Printf("ERR: %v\n", err)

	// delete
	//num, err = o.Delete(&u)
	//fmt.Printf("NUM: %d, ERR: %v\n", num, err)
	beego.Run()
}

//import (
//	"fmt"
//	"github.com/astaxie/beego/orm"
//	_ "github.com/go-sql-driver/mysql" // import your used driver
//)
//
//// Model Struct
//type User struct {
//	Id   int
//	Name string `orm:"size(100)"`
//}
//
//func init() {
//	// set default database
//	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/kcp?charset=utf8", 30)
//
//	// register model
//	orm.RegisterModel(new(User))
//
//	// create table
//	orm.RunSyncdb("default", false, true)
//}
//
//func main() {
//	o := orm.NewOrm()
//
//	user := User{Name: "slene"}
//
//	// insert
//	id, err := o.Insert(&user)
//	fmt.Printf("ID: %d, ERR: %v\n", id, err)
//
//	// update
//	user.Name = "astaxie"
//	num, err := o.Update(&user)
//	fmt.Printf("NUM: %d, ERR: %v\n", num, err)
//
//	// read one
//	u := User{Id: user.Id}
//	err = o.Read(&u)
//	fmt.Printf("ERR: %v\n", err)
//
//	// delete
//	num, err = o.Delete(&u)
//	fmt.Printf("NUM: %d, ERR: %v\n", num, err)
//}
