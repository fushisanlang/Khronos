/*
@Time : 2021/8/3 3:34 下午
@Author : fushisanlang
@File : user.go
@Software: GoLand
*/
package dao

import (
	"Khronos/app/model"
	"fmt"
	"github.com/gogf/gf/frame/g"
)

//INSERT INTO `a_game`.`user`(`user`, `pass`, `email`, ) VALUES ();
func UserAdd(user *model.AddUser) {
	fmt.Println(user)
	result, err := g.DB().Model("user").Data(user).Insert()
	fmt.Println(result)
	fmt.Println(err)
}

//select count(*) from user where user = username
func UserNameGet(username string) int {
	a, _ := g.DB().Model("user").Where("username", username).Count()
	return a
}

//select pass from user where user = username
func PassGetByName(username string) string {
	passStr := (*model.PassStr)(nil)
	err := g.DB().Model("user").Where("username", username).Scan(&passStr)
	if err != nil {
		return ""
	} else if passStr == nil {
		return ""
	} else {
		return passStr.PassStr
	}
}

//select uid from user where user = username
func UidGetByName(username string) int {
	fmt.Println(username)
	uid := (*model.UserId)(nil)
	err := g.DB().Model("user").Where("username", username).Scan(&uid)
	fmt.Println(uid)
	if err != nil {
		return -1
	} else if uid == nil {
		return -1
	} else {
		return uid.UserId
	}
}
