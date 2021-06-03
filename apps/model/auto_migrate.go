package model

import (
	. "github.com/LangPham/mila_go/apps/repo"
)

func init() {
	_ = Repo.AutoMigrate(&Account{}, &Credential{}, &Tag{}, &Article{})
}
