package action

import (
	. "github.com/LangPham/mila_go/apps/model"
	. "github.com/LangPham/mila_go/apps/repo"
)

func ListTag() (tags []Tag) {
	Repo.Find(&tags)
	return
}

//func CreateTag(attr []byte) (exchange Exchange) {
//	tag := new(Tag)
//	exchange = tag.Change(attr)
//	if exchange.Valid {
//		Repo.Create(&tag)
//		exchange.ResultID = string(tag.ID)
//	}
//	return
//}

//func GetTag(id string) (result Tag) {
//	Repo.First(&result, id)
//	return
//}
//
//func UpdateTag(tag Tag, attrs interface{}) (error bool, result Change) {
//	var changeset Change
//
//	changeset = tag.Change(attrs)
//	tag = changeset.Data.(Tag)
//
//	if changeset.Valid {
//		Repo.Save(&tag)
//		changeset.IdResult = strconv.Itoa(int(tag.ID))
//
//		result = changeset
//		error = true
//	} else {
//		result = changeset
//		error = false
//	}
//	return
//}
//
//func DeleteTag(id string) (result Tag) {
//	Repo.Delete(&Tag{}, id)
//	return
//}