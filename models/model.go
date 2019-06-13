package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)
type User struct {
	Id int
	Name string
	Passwd string
}
func init(){
	//1.连接数据库
	orm.RegisterDataBase("default","mysql","root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	//2.注册表
	orm.RegisterModel(new(User))
	//3.生成表
	//1.数据库别名
	//2.是否强制更新
	//3.创建表过程是否可见
	orm.RunSyncdb("default",false,true)
}
