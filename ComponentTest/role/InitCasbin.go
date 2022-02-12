package role

import (
	"fmt"
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

var e *casbin.Enforcer

func InitCasbin() {
	dsName := "root:admin@tcp(127.0.0.1:3306)/gf_test"
	a, _ := gormadapter.NewAdapter("mysql", dsName, true)
	newE, err := casbin.NewEnforcer("role/rbac_model.conf", a)
	if err != nil {
		fmt.Println(err.Error())
	}
	e = newE
	e.LoadPolicy()
}

func GetCasbinEnforcer() *casbin.Enforcer {
	return e
}
