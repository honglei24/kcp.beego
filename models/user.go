package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type UserStatus int

const (
	Normal UserStatus=iota
	Deleted
	Freeze
)

const UserTableName="user"

type User struct {
	UserId			int `orm:"pk"`
	RoleId			int
	Username    	string `orm:"size(50)", valid:"Required"`
	Password     	string	`orm:"size(50)", valid:"Required"`
	Mobile     		string	`orm:"size(15)"`
	Email     		string	`orm:"size(50)"`
	Status 			UserStatus
	Cluster 		[]*Cluster	`orm:"reverse(many)"`
	CreateTime 		time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime		time.Time `orm:"auto_now;type(datetime)"`
}

func UserAdd(u *User) (int64, error) {
	return orm.NewOrm().Insert(u)
}

func UserGetList() ([]*User, int64) {
	list := make([]*User, 0)
	query := orm.NewOrm().QueryTable(UserTableName)
	total, _ := query.Count()
	query.OrderBy("user_id").All(&list)
	return list, total
}

func UserGetById(id int) (*User, error) {
	r := new(User)
	err := orm.NewOrm().QueryTable(UserTableName).Filter("user_id", id).One(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (a *User) TableName() string {
	return UserTableName
}

func (a *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(a, fields...); err != nil {
		return err
	}
	return nil
}

//func (u *User) Valid(v *validation.Validation) {
//	if strings.Index(u.Username, "admin") != -1 {
//		// 通过 SetError 设置 Name 的错误信息，HasErrors 将会返回 true
//		v.SetError("Name", "名称里不能含有 admin")
//	}
//}