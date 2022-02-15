package auth

import (
	"GFBackend/config"
	"GFBackend/logger"
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
	appConfig := config.AppConfig

	dsName := appConfig.Database.Username + ":" +
		appConfig.Database.Password + "@tcp(" +
		appConfig.Database.IP + ")/" +
		appConfig.Database.DB1
	a, _ := gormadapter.NewAdapter("mysql", dsName, true)
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

}
