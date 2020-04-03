package models

import "time"

type Right struct {
	RightId         int  `orm:"pk"`
	RightName   string	 `orm:"size(50)"`
	Detail     string
	Status     int
	CreateId   int
	UpdateId   int
	CreateTime 		time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime		time.Time `orm:"auto_now;type(datetime)"`
}