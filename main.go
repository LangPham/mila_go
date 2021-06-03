package main

import (
	"github.com/LangPham/mila_go/apps_web"
)

func main() {
	//a :=config.Config.Get("views")
	//aon.Dump(apps_web.Cas, "CAS")
	//enforce, err := apps_web.Cas.Enforce("member", "/member/1", "*")
	//if err != nil {
	//	return
	//}
	//fmt.Println(enforce)
	//enc := util.MilaEncrypt("cai gi day")
	//util.Dump(enc, "ENC")
	//dec := util.MilaDecrypt(enc)
	//util.Dump(dec, "DEC")
	apps_web.AppsWeb()
}

