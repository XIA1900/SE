package service

import (
	"GFBackend/middleware/auth"
	"GFBackend/model/dao"
	"GFBackend/utils"
)

func Register(username, password string) {
	salt := utils.GetRandomString(6)
	newUser := dao.User{
		Username: username,
		Password: utils.EncodeInMD5(password + salt),
		Salt:     salt,
	}
	dao.CreateUser(newUser)
	auth.CasbinEnforcer.AddPolicy(username, "regular")
}

func Login() {

}

func Logout() {

}

func UpdatePassword() {

}

func Delete() {

}
