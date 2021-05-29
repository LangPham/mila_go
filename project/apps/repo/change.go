package repo

import (
	dbug "github.com/LangPham/mila_debug"
	"github.com/emirpasic/gods/maps/hashmap"
	"github.com/gofiber/fiber/v2"
	"reflect"
)

type Exchange struct {
	Valid    bool
	Error    *hashmap.Map
	Change   *hashmap.Map
	Data     interface{}
	Action   string
	DataType string
	ResultID string
}

func NewExchange(dataIn interface{}) *Exchange {
	switch v := dataIn.(type) {
	case string:
		return &Exchange{
			Valid:    false,
			Error:    hashmap.New(),
			Change:   hashmap.New(),
			Action:   "new",
			DataType: v,
		}
	default:
		dataType := reflect.TypeOf(dataIn).Name()
		dbug.Dump(dataType, "NEWEXCHANGE")
		return &Exchange{
			Valid:    false,
			Error:    hashmap.New(),
			Change:   hashmap.New(),
			Data:     dataIn,
			Action:   "edit",
			DataType: dataType,
		}
	}
}

func attToTreemap(c *fiber.Ctx, properties []string) *hashmap.Map {
	m := hashmap.New()
	for _, value := range properties {
		m.Put(value, c.FormValue(value))
	}
	return m
}

func Cast(modelIn interface{}, attr *hashmap.Map, properties ...string) (exchange Exchange) {

	//dataType := reflect.TypeOf(modelIn).Name()
	modelInType := reflect.ValueOf(&modelIn).Elem()
	newModel := reflect.New(modelInType.Elem().Type()).Elem()
	dbug.Dump(newModel, "newModel")

	//mAttr := attToTreemap(c, properties)
	//dbug.Dump(mAttr, "Treemap")
	//for i := 0; i < len(properties); i++ {
	//	value, err := mAttr.Get(properties[i])
	//	//dbug.Dump(value, "Value: "+properties[i])
	//	//dbug.Dump(newModel.FieldByName(properties[i]).Kind().String(), "Type :"+properties[i])
	//	if err && value != nil {
	//
	//		switch newModel.FieldByName(properties[i]).Kind().String() {
	//		case "string":
	//			newModel.FieldByName(properties[i]).SetString(value.(string))
	//		case "int64":
	//			//Todo: err
	//			vint, _ := strconv.Atoi(value.(string))
	//			newModel.FieldByName(properties[i]).SetInt(int64(vint))
	//		default:
	//			//dbug.Dump(newModel.FieldByName(properties[i]).Kind().String(), "Undefined")
	//			//newModel.FieldByName(properties[i]).SetString("Undefined")
	//		}
	//	}
	//
	//}
	//modelInType.Set(newModel)
	//exchange.Data = modelIn
	//action := "new"
	//value, err := mAttr.Get("_METHOD")
	//if err && value != nil {
	//	//dbug.Dump(value, "METHOD")
	//	action += value.(string)
	//	switch strings.ToLower(value.(string)) {
	//	case "post":
	//		action = "insert"
	//	case "put":
	//		action = "update"
	//	default:
	//		action = "new"
	//	}
	//}
	//exchange.Action = action
	//exchange.DataType = dataType
	//exchange.Change = mAttr
	return
}

func (exchange *Exchange) ValidateRequired(properties ...string) {
	exchange.Valid = true
	exError := hashmap.New()
	for i := 0; i < len(properties); i++ {
		value, exist := exchange.Change.Get(properties[i])
		if !exist || value == "" {
			exError.Put(properties[i], properties[i]+" is required!")
			exchange.Valid = false
		}
	}
	exchange.Error = exError
}
