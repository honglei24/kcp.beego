package models

type UserStatus int

const (
	Normal UserStatus=iota
	Deleted
	Freeze
)
type User struct {
	Id				int
	RoleId			int
	Username    	string
	Password     	string
	Status 			UserStatus
}

type Cluster struct {
	Id          	int
	UserId         	int
	Version			string
	KubeConfig		string
	Shared			int
	CapacityCpu		int
	CapacityMem		int64
	Status 			string
}

//type Profile struct {
//	Id          int
//	Age         int16
//	User        *User   `orm:"reverse(one)"` // 设置一对一反向关系(可选)
//}
//
//type Post struct {
//	Id    int
//	Title string
//	User  *User  `orm:"rel(fk)"`    //设置一对多关系
//	Tags  []*Tag `orm:"rel(m2m)"`
//}
//
//type Tag struct {
//	Id    int
//	Name  string
//	Posts []*Post `orm:"reverse(many)"` //设置多对多反向关系
//}
