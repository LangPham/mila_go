package repo

import (
	"fmt"
	"github.com/emirpasic/gods/maps/hashmap"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
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

func NewExchange(dataIn interface{}) *Exchange {
	switch v := dataIn.(type) {
	case string:
		return &Exchange{
			Valid:    false,
			Error:    hashmap.New(),
			Change:   hashmap.New(),
			Request:  "new",
			DataType: v,
			Validate: validator.New(),
		}
	default:
		dataType := reflect.TypeOf(dataIn).Name()

		return &Exchange{
			Valid:    false,
			Error:    hashmap.New(),
			Change:   hashmap.New(),
			Data:     dataIn,
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
	exchange.Valid = mError == nil
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
	//aon.Dump(exchange, "VALI")
	validate := validator.New()

	// returns nil or ValidationErrors ( []FieldError )
	err := validate.Struct(exchange.Data)

	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}

		for _, err := range err.(validator.ValidationErrors) {

			switch {
			case contains([]string{"email"}, err.Tag()):
				exchange.Error.Put(err.Field(), "not " + err.Tag())
			case contains([]string{"required"}, err.Tag()):
				exchange.Error.Put(err.Field(), err.Tag())
			case contains([]string{"eq","gt","gte","lt","lte","ne" }, err.Tag()):
				exchange.Error.Put(err.Field(),"not "+ err.Tag() + " " + err.Param())
			default:
				exchange.Error.Put(err.Field(),"not "+ err.Tag() + " " + err.Param())
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
		exchange.Valid = exchange.Error == nil
		return
	}
}

