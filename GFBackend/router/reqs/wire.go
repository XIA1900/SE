//go:build wireinject
// +build wireinject

package reqs

import (
	"GFBackend/controller"
	"github.com/google/wire"
)

func InitializeUserManageController() (*controller.UserManageController, error) {
	panic(wire.Build(controller.UserManageControllerSet))
}

func InitializeCommunityManageController() (*controller.CommunityManageController, error) {
	panic(wire.Build(controller.CommunityManageSet))
}

func InitializeFileManageController() (*controller.FileManageController, error) {
	panic(wire.Build(controller.FileManageControllerSet))
}
