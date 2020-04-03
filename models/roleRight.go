package models

import "time"

type RoleRight struct {
	RightId         int  `orm:"pk"`
	RoleId     		int
	CreateTime 		time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime		time.Time `orm:"auto_now;type(datetime)"`
}