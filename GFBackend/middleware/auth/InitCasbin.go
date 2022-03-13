package auth

import (
	"GFBackend/config"
	"GFBackend/logger"
	"GFBackend/model"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

var CasbinEnforcer *casbin.Enforcer

func InitCasbin() {
	a, _ := gormadapter.NewAdapterByDB(model.NewDB())
	e, err := casbin.NewEnforcer("middleware/auth/rbac_model.conf", a)
	CasbinEnforcer = e
	if err != nil {
		logger.AppLogger.Error(err.Error())
		panic(err)
	}
	err = CasbinEnforcer.LoadPolicy()
	if err != nil {
		logger.AppLogger.Error(err.Error())
		panic(err)
	}

	addInitialPolicy()
}

func addInitialPolicy() {
	basePath := config.AppConfig.Server.BasePath

	// regular

	// /user/...
	CasbinEnforcer.AddPolicy("regular", basePath+"/user/logout", "POST")
	CasbinEnforcer.AddPolicy("regular", basePath+"/user/password", "POST")
	CasbinEnforcer.AddPolicy("regular", basePath+"/user/update", "POST")
	CasbinEnforcer.AddPolicy("regular", basePath+"/user/follow", "POST")
	CasbinEnforcer.AddPolicy("regular", basePath+"/user/unfollow", "POST")
	CasbinEnforcer.AddPolicy("regular", basePath+"/user/followers", "POST")
	CasbinEnforcer.AddPolicy("regular", basePath+"/user/followees", "POST")

	// /community/...
	CasbinEnforcer.AddPolicy("regular", basePath+"/community/create", "POST")

	// /file/...
	CasbinEnforcer.AddPolicy("regular", basePath+"/file/upload", "POST")
	CasbinEnforcer.AddPolicy("regular", basePath+"/file/download", "POST")
	CasbinEnforcer.AddPolicy("regular", basePath+"/file/delete", "POST")
	CasbinEnforcer.AddPolicy("regular", basePath+"/file/scan", "POST")
	CasbinEnforcer.AddPolicy("regular", basePath+"/file/space/info", "POST")

	// admin
	CasbinEnforcer.AddGroupingPolicy("admin", "regular") // admin extends regular
	CasbinEnforcer.AddPolicy("admin", basePath+"/user/admin/register", "POST")
	CasbinEnforcer.AddPolicy("admin", basePath+"/user/admin/delete", "POST")
	CasbinEnforcer.AddPolicy("admin", basePath+"/file/space/update", "POST")
	CasbinEnforcer.AddPolicy("admin", basePath+"/articletype/create", "POST")

	// default admin user
	CasbinEnforcer.AddGroupingPolicy("boss", "admin")
}
