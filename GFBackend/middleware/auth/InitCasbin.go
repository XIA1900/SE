package auth

import (
	"GFBackend/config"
	"GFBackend/logger"
	"GFBackend/model"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

type CasbinRule struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Ptype string `gorm:"size:512;uniqueIndex:unique_index"`
	V0    string `gorm:"size:512;uniqueIndex:unique_index"`
	V1    string `gorm:"size:512;uniqueIndex:unique_index"`
	V2    string `gorm:"size:512;uniqueIndex:unique_index"`
	V3    string `gorm:"size:512;uniqueIndex:unique_index"`
	V4    string `gorm:"size:512;uniqueIndex:unique_index"`
	V5    string `gorm:"size:512;uniqueIndex:unique_index"`
}

var CasbinEnforcer *casbin.Enforcer

func InitCasbin() {
	// appConfig := config.AppConfig
	//dsName := appConfig.Database.Username + ":" +
	//	appConfig.Database.Password + "@tcp(" +
	//	appConfig.Database.IP + ")/" +
	//	appConfig.Database.DB1
	//a, _ := gormadapter.NewAdapter("mysql", dsName, true)
	a, _ := gormadapter.NewAdapterByDB(model.DB)
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
	CasbinEnforcer.AddPolicy("regular", basePath+"/user/logout", "POST")
	CasbinEnforcer.AddPolicy("regular", basePath+"/user/password", "POST")
	CasbinEnforcer.AddPolicy("regular", basePath+"/user/update", "POST")

	// admin
	CasbinEnforcer.AddGroupingPolicy("admin", "regular") // admin extends regular
	CasbinEnforcer.AddPolicy("admin", basePath+"/user/admin/register", "POST")
	CasbinEnforcer.AddPolicy("admin", basePath+"/user/admin/delete", "POST")
}
