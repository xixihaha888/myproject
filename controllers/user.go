package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"myproject/models"
)

type UserController struct{
	beego.Controller
}
func (this*UserController)ShowRegister(){
	this.TplName = "register.html"
}
func (this*UserController)HandleRegister(){
	userName := this.GetString("userName")
	passwd := this.GetString("passwd")
	if userName == "" || passwd == ""{
		beego.Info("数据数据不完整，请重新输入！")
		this.TplName = "register.html"
		return
	}
	//获取orm对象
	o := orm.NewOrm()
	//获取要插入的数据对象
	var user models.User
	//给对象赋值
	user.Name = userName
	user.Passwd = passwd
	//把数据插入到数据库
	if _,err := o.Insert(&user);err != nil{
		beego.Info("注册失败，请更换用户名再次注册!")
		this.TplName = "register.html"
		return
	}
	this.Ctx.WriteString("注册成功!")

}
//显示登陆界面
func(this*UserController)ShowLogin(){
	this.TplName = "login.html"
}
//处理登陆业务
func(this*UserController)HandleLogin(){
	//获取前端传递的数据
	userName := this.GetString("userName")
	passwd := this.GetString("passwd")
	//对数据进行校验
	if userName == "" || passwd == ""{
		beego.Info("数据数据不完整，请重新输入！")
		this.TplName = "login.html"
		return
	}
	//查询数据库，判断用户名和密码是否正确
	//获取orm对象
	o := orm.NewOrm()
	//获取要插入的数据对象
	var user models.User
	//给对象赋值
	user.Name = userName
	//根据用户名查询
	if err := o.Read(&user,"Name");err != nil{
		beego.Info("用户名错误，请重新输入！")
		this.TplName = "login.html"
		return
	}
	if user.Passwd != passwd{
		beego.Info("密码错误，请重新输入！")
		this.TplName = "login.html"
		return
	}

	//返回提示信息
	this.Ctx.WriteString("登陆成功!")
}


