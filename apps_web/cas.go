package apps_web

import (
	"github.com/LangPham/mila/apps/aon"
	"github.com/casbin/casbin/v2"
)

var Cas *casbin.Enforcer
func init() {
	e, _ := casbin.NewEnforcer("config/rbac_model.conf", "config/rbac_policy.csv")
	e.LoadModel()
	e.LoadPolicy()
	Cas = e

	enforce, err := e.Enforce("member", "/member/*", "*")
	if err != nil {
		return 
	}
	aon.Dump(enforce)

}