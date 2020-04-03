/*
* @Author: honglei
* @Date: 2020/4/1 17:00
 */
package sysinit

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"  // 引入orm的包
	_ "github.com/go-sql-driver/mysql"   // 引入MySQL包
	"kcp/models"
)

func initDB() {
	dbAlias := beego.AppConfig.String("db_alias")
	dbName := beego.AppConfig.String("db_name")
	dbUser := beego.AppConfig.String("db_user")
	dbPwd := beego.AppConfig.String("db_pwd")
	dbHost := beego.AppConfig.String("db_host")
	dbPort := beego.AppConfig.String("db_port")
	dbCharset := beego.AppConfig.String("db_charset")
	maxIdle, err := beego.AppConfig.Int("max_idle")
	if err != nil {
		logs.Error("max_idle must be an integer.")
	}
	maxConn, err := beego.AppConfig.Int("max_conn")
	if err != nil {
		logs.Error("max_conn must be an integer.")
	}

	// set default database
	orm.RegisterDataBase(dbAlias, "mysql", dbUser+":"+dbPwd+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?"+dbCharset, maxIdle, maxConn)

	// register model
	// TODO
	orm.RegisterModel(new(models.User), new(models.Cluster), new(models.Role), new(models.Right), new(models.RoleRight))

	// 如果是开发模式， 则显示命令信息
	isDev := (beego.AppConfig.String("runmode") == "dev")
	// create table
	err = orm.RunSyncdb("default", false, isDev)

	if err != nil {
		logs.Informational("[orm] Create table err : ", err)
	}
	if isDev {
		orm.Debug = isDev
	}
}

