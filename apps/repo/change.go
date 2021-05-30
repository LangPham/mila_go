package repo

import (
	"github.com/LangPham/mila_go/apps/aon"
	"github.com/emirpasic/gods/maps/hashmap"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/huandu/xstrings"
	"reflect"
	"strconv"
	"strings"
)

type Exchange struct {
	Valid    bool
	Error    *hashmap.Map
	Change   *hashmap.Map
	Data     interface{}
	Request  string
	DataType string
	ResultID string
	Validate *validator.Validate
}

func NewExchange(modelIn interface{}) *Exchange {
	switch v := modelIn.(type) {
	case string:
		//aon.Dump(modelIn, "NEW EXCHANGE string")
		return &Exchange{
			Valid:    false,
			Error:    hashmap.New(),
			Change:   hashmap.New(),
			Request:  "new",
			DataType: v,
			Validate: validator.New(),
		}
	default:
		dataType := reflect.TypeOf(modelIn).Name()
		//aon.Dump(modelIn, "Edit EXCHANGE data")
		change := hashmap.New()
		modelInType := reflect.ValueOf(&modelIn).Elem()
		for i := 0; i < modelInType.Elem().NumField(); i++ {
			fieldName := modelInType.Elem().Type().Field(i).Tag.Get("cast")
			if fieldName != "" {
				change.Put(fieldName, modelInType.Elem().Field(i).Interface())
			}
			//aon.Dump(modelInType.Elem().Field(i).Interface(), "VALUE")
		}

		return &Exchange{
			Valid:    false,
			Error:    hashmap.New(),
			Change:   change,
			Data:     modelIn,
			Request:  "edit",
			DataType: dataType,
			Validate: validator.New(),
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

func Cast(modelIn interface{}, c *fiber.Ctx) (exchange Exchange) {
	//aon.Dump(modelIn, "MODELIN")
	//aon.Dump(c, "ATTR")
	change := hashmap.New()
	mError := hashmap.New()
	//aon.Dump(c.FormValue("UserName"), "ATTR")
	dataType := reflect.TypeOf(modelIn).Name()
	modelInType := reflect.ValueOf(&modelIn).Elem()
	//aon.Dump(modelInType.Elem().Type(), "modelInType")
	newModel := reflect.New(modelInType.Elem().Type()).Elem()
	//aon.Dump(newModel, "NEW MODEL")
	//aon.Dump(newModel.Type(), "NUM Field")
	for i := 0; i < newModel.NumField(); i++ {
		//aon.Dump(newModel.Type().Field(i).Name, "NAME")
		//aon.Dump(newModel.Type().Field(i).Tag, "Tag")
		//aon.Dump(newModel.Type().Field(i).Tag.Get("cast"), "CASTTag")
		//aon.Dump(newModel.Type().Field(i).Type, "TYPE")
		//aon.Dump(newModel.Field(i).Interface(), "INT")
		//value := c.FormValue()
		//strings.Split(newModel.Type().Field(i).Tag.(string), " ")
		fieldName := newModel.Type().Field(i).Tag.Get("cast")
		if fieldName != "" {
			fieldValue := c.FormValue(fieldName)
			change.Put(fieldName, fieldValue)
			//aon.Dump(newModel.Type().Field(i).Type.String(), "TYPE")
			switch newModel.Type().Field(i).Type.String() {
			case "string":
				newModel.Field(i).SetString(fieldValue)
			case "int":
				//aon.Dump(fieldValue, "VAL INT")
				valueInt, err := strconv.Atoi(fieldValue)
				if err != nil {
					mError.Put(fieldName, "Please input number!")
				} else {
					newModel.Field(i).SetInt(int64(valueInt))
				}
			}
		}
	}

	modelInType.Set(newModel)
	exchange.Data = modelIn

	value := c.FormValue("_METHOD")
	userReq := ""
	switch strings.ToLower(value) {
	case "post":
		userReq = "insert"
	case "put":
		userReq = "update"
	default:
		userReq = "new"
	}
	exchange.Request = userReq
	exchange.DataType = dataType
	exchange.Change = change
	exchange.Error = mError
	exchange.Valid = mError.Empty()
	return
}

func contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func (exchange *Exchange) ValidateModel(properties ...string) {

	validate := validator.New()

	err := validate.Struct(exchange.Data)
	if err != nil {
		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		//if _, ok := err.(*validator.InvalidValidationError); ok {
		//	fmt.Println(err)
		//	return
		//}

		for _, err := range err.(validator.ValidationErrors) {

			switch {
			case contains([]string{"email"}, err.Tag()):
				exchange.putError(err.Field(), "Please input email!")
			case contains([]string{"required"}, err.Tag()):
				exchange.putError(err.Field(), "Can't blank!")
			case contains([]string{"eq", "gt", "gte", "lt", "lte", "ne"}, err.Tag()):
				exchange.putError(err.Field(), "not "+err.Tag()+" "+err.Param())
			default:
				exchange.putError(err.Field(), "not "+err.Tag()+" "+err.Param())
				//fmt.Println(err.Namespace())
				//fmt.Println(err.Field())
				//fmt.Println(err.StructNamespace())
				//fmt.Println(err.StructField())
				//fmt.Println(err.Tag())
				//fmt.Println(err.ActualTag())
				//fmt.Println(err.Kind())
				//fmt.Println(err.Type())
				//fmt.Println(err.Value())
				//fmt.Println(err.Param())
				//fmt.Println()
				//fmt.Println(err)
			}
		}
		exchange.Valid = exchange.Error.Empty()
		return
	}
}

func (exchange *Exchange) putError(field string, err string) {
	key := xstrings.ToSnakeCase(field)
	val, has := exchange.Error.Get(key)
	if has {
		exchange.Error.Put(key, val.(string)+", "+err)
	} else {
		exchange.Error.Put(key, err)
	}
}

func (exchange *Exchange) PutField(field string, value interface{}) {
	//exchange.Data
	rData := reflect.ValueOf(&exchange.Data).Elem()
	//aon.Dump(rData.Elem(), "ELEM")
	oldModel := rData.Elem()
	aon.Dump(oldModel.Field(1))
	newModel := reflect.New(rData.Elem().Type()).Elem()

	aon.Dump(newModel, "new")
	//aon.Dump(newModel.NumField(), "so field")
	for i := 0; i < newModel.NumField(); i++ {
		if newModel.Type().Field(i).Name == field {
			// Put data
			switch newModel.Type().Field(i).Type.String() {
			case "string":
				newModel.Field(i).SetString(value.(string))
			case "int":
				newModel.Field(i).SetInt(value.(int64))
			}
		} else {
			newModel.Field(i).Set(oldModel.Field(i))
		}

	}
	rData.Set(newModel)
	//exchange.Data =
	//modelInType.NumField()

	//aon.Dump(reflect.ValueOf(&modelInType), "PUT FIELD")
}
