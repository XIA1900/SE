package test

import (
	"GFBackend/model/dao"
	"GFBackend/service"
	"testing"
)

func TestGetCommunityByName(t *testing.T) {
	communityManageService := service.NewCommunityManageService(dao.NewCommunityDAO(), dao.NewUserDAO())
	community, user, _ := communityManageService.GetCommunityByName("test")
	if community.Name != "test" {
		t.Error("GetCommunityByName error")
	}
	if user.Nickname != "test" {
		t.Error("GetCommunityByName error")
	}
}
