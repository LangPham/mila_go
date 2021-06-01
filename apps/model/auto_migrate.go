package model

import (
	. "github.com/LangPham/mila/apps/repo"
)

func init() {
	_ = Repo.AutoMigrate(&Account{}, &Credential{}, &Tag{}, &Article{})
}
