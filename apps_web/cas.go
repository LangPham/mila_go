package apps_web

import (
	"github.com/casbin/casbin/v2"
)

var Cas *casbin.Enforcer

func init() {
	e, _ := casbin.NewEnforcer("config/rbac_model.conf", "config/rbac_policy.csv")
	e.LoadModel()
	e.LoadPolicy()


	Cas = e
	//
	//enforce, err := Cas.Enforce("member", "/member/1/aa", "read")
	//util.Dump(err)
	//util.Dump(enforce)

}