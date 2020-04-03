package models

import "time"

type Cluster struct {
	ClusterId       int `orm:"pk"`
	UserId        	*User	`orm:"rel(fk);null;on_delete(do_nothing)"`
	Version			string	`orm:"size(20)"`
	KubeConfig		string	`orm:"type(text)"`
	Shared			bool
	CapacityCpu		int
	CapacityMem		int64
	Status 			string
	CreateTime 		time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime		time.Time `orm:"auto_now;type(datetime)"`
}
