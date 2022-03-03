package test

import (
	"GFBackend/model/dao"
	"GFBackend/service"
	"testing"
)

var userManageService = service.NewUserManageService(dao.NewUserDAO(), dao.NewFollowDAO(), dao.NewSpaceDAO())

func TestUserRegister(t *testing.T) {
	err := userManageService.Register("bird", "007", false)
	if err != nil {
		t.Error(err.Error())
		return
	}
}
