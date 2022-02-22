//go:build wireinject
// +build wireinject

package router

import (
	"GFBackend/controller"
	"github.com/google/wire"
)

func InitializeUserManageController() (*controller.UserManageController, error) {
	panic(wire.Build(controller.UserManageSet))
}

func InitializeCommunityManageController() (*controller.CommunityManageController, error) {
	panic(wire.Build(controller.CommunityManageSet))
}
