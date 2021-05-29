package mila_view

import (
	"github.com/aymerick/raymond"
	"path/filepath"
)

func init() {
	raymond.RegisterHelper("include", func(val1 string) raymond.SafeString {
		return raymond.SafeString(Include(val1))
	})

	raymond.RegisterHelper("render_with_context", func(file string, options *raymond.Options) raymond.SafeString {
		ctx := options.Ctx()
		return raymond.SafeString(renderWithContext(file, ctx))
	})

	raymond.RegisterHelper("form_for", func(changeset interface{}, action string, options *raymond.Options) raymond.SafeString {
		return raymond.SafeString(FormFor(changeset, action, options))
	})

	raymond.RegisterHelper("link", func(to string, content string, options *raymond.Options) raymond.SafeString {
		method := ""
		methodLink := options.HashStr("method")
		if len(methodLink) > 0 {
			method = `data-method="` + methodLink + `" `
			if methodLink == "DELETE" {
				method = method + `data-confirm="Are you sure?"`
			}
		}
		ctxLink := options.Hash()
		tpl := raymond.MustParse(to)

		resultIn := tpl.MustExec(ctxLink)
		resultBefore := `<a href="` + resultIn + `" ` + method + ` >`
		resultAfter := `</a>`
		result := resultBefore + content + resultAfter

		return raymond.SafeString(result)
	})

	//raymond.RegisterHelper("tag_input", func(changeset interface{}, fieldName string, options *raymond.Options) raymond.SafeString {
	//	modelName := ""
	//	idName := ""
	//	errStr := ""
	//
	//	switch changeset.(type) {
	//	case string:
	//
	//		modelName = changeset.(string)
	//	default:
	//		change := changeset.(repo.Change)
	//		modelName = change.DataType
	//		mapError := change.Error
	//		idName = modelName + "_" + strings.ToLower(fieldName)
	//		errTest, _ := mapError.Get(idName)
	//		if errTest != nil {
	//			errStr = `<p class="error">` + fieldName + `: ` + errTest.(string) + `</p>`
	//		}
	//	}
	//
	//	value := options.HashStr("value")
	//	class := options.HashStr("class")
	//	vType := "text"
	//	if len(options.HashStr("type")) > 0 {
	//		vType = options.HashStr("type")
	//	}
	//
	//	proId := `id="` + modelName + `_` + strings.ToLower(fieldName) + `" `
	//	proName := `name="` + modelName + `[` + fieldName + `]" `
	//	proType := `type="` + vType + `" `
	//	proValue := `value="` + value + `"`
	//	proClass := `class="` + class + `"`
	//	result := "<input " + proId + proName + proType + proValue + proClass + `>` + errStr
	//	return raymond.SafeString(result)
	//})
	//
	//raymond.RegisterHelper("tag_select", func(changeset interface{}, fieldName string, options *raymond.Options) raymond.SafeString {
	//	modelName := ""
	//	idName := ""
	//	errStr := ""
	//
	//	switch changeset.(type) {
	//	case string:
	//		modelName = changeset.(string)
	//	default:
	//		change := changeset.(repo.Change)
	//		modelName = change.DataType
	//		mapError := change.Error
	//		idName = modelName + "_" + strings.ToLower(fieldName)
	//		errTest, _ := mapError.Get(idName)
	//		if errTest != nil {
	//			errStr = `<p class="error">` + fieldName + `: ` + errTest.(string) + `</p>`
	//		}
	//	}
	//
	//	data := options.HashProp("data")
	//	value := options.HashStr("value")
	//	text := options.HashStr("text")
	//	prompt := options.HashStr("prompt")
	//	selected := options.HashStr("selected")
	//
	//
	//	var optionHTML strings.Builder
	//	if len(prompt) >0 {
	//		optionStr := `<option value="">`+ prompt +`</option>`
	//		optionHTML.WriteString(optionStr)
	//	}
	//	lenArr := gjson.GetBytes(data.([]byte), "#")
	//
	//	var i int64 = 0
	//	for ; i <lenArr.Int(); i++ {
	//		valueResult := gjson.GetBytes(data.([]byte), strconv.FormatInt(i, 10)+"."+value)
	//		textResult := gjson.GetBytes(data.([]byte), strconv.FormatInt(i, 10)+"."+text)
	//
	//		strValue := ""
	//		switch valueResult.Type.String() {
	//		case "Number":
	//			strValue = fmt.Sprintf("%f", valueResult.Num)
	//			slice := strings.Split(strValue, ".")
	//			if len(strings.TrimRight(slice[1], "0")) == 0 {
	//				strValue = fmt.Sprintf("%.0f", valueResult.Num)
	//			}
	//		case "String":
	//			strValue = valueResult.Str
	//
	//		default:
	//			strValue = ""
	//		}
	//		textSelected := ""
	//		if strValue == selected {
	//			textSelected = "selected"
	//		}
	//		optionStr := `<option value="`+ strValue +`" `+ textSelected +`>`+textResult.Str +`</option>`
	//		optionHTML.WriteString(optionStr)
	//	}
	//
	//	class := options.HashStr("class")
	//
	//	proId := `id="` + modelName + `_` + strings.ToLower(fieldName) + `" `
	//	proName := `name="` + modelName + `[` + fieldName + `]" `
	//
	//	proClass := `class="` + class + `"`
	//	result := "<select " + proId + proName +  proClass + `>`+ optionHTML.String() +`</select>` + errStr
	//	return raymond.SafeString(result)
	//})
	//
	//raymond.RegisterHelper("tag_textarea", func(changeset interface{}, fieldName string, options *raymond.Options) raymond.SafeString {
	//	modelName := ""
	//	idName := ""
	//	errStr := ""
	//	switch changeset.(type) {
	//	case string:
	//		modelName = changeset.(string)
	//	default:
	//		change := changeset.(repo.Change)
	//		modelName = change.DataType
	//		mapError := change.Error
	//		idName = modelName + "_" + strings.ToLower(fieldName)
	//		errTest, _ := mapError.Get(idName)
	//		if errTest != nil {
	//			errStr = `<p class="error">` + fieldName + `: ` + errTest.(string) + `</p>`
	//		}
	//	}
	//
	//	value := options.HashStr("value")
	//	class := options.HashStr("class")
	//	vType := "text"
	//	if len(options.HashStr("type")) > 0 {
	//		vType = options.HashStr("type")
	//	}
	//
	//	proId := `id="` + modelName + `_` + strings.ToLower(fieldName) + `" `
	//	proName := `name="` + modelName + `[` + fieldName + `]" `
	//	proType := `type="` + vType + `" `
	//	//proValue := `value="` + value + `"`
	//	proClass := `class="` + class + `"`
	//	result := "<textarea " + proId + proName + proType + proClass + `>` + value + `</textarea>` + errStr
	//	return raymond.SafeString(result)
	//})

}

func Include(file string) (result string) {
	result = ""
	fullFile := file + ".hgo"
	fileTpl := filepath.Join("theta_web/templates", fullFile)
	tpl, e := raymond.ParseFile(fileTpl)
	if e != nil {
		return
	}
	result = tpl.MustExec(nil)
	return
}

func renderWithContext(file string, ctx interface{}) (result string) {

	result = ""
	fullFile := file + ".hgo"
	fileTpl := filepath.Join("theta_web/templates", fullFile)
	tpl, e := raymond.ParseFile(fileTpl)
	if e != nil {
		return
	}

	result = tpl.MustExec(ctx)
	return
}

func FormFor(changeset interface{}, action string, options *raymond.Options) (result string) {
	enctype := options.HashStr("enctype")
	enctypeAtt := ""
	if len(enctype) > 0 {
		enctypeAtt = `enctype="` + enctype + `"`
	}

	method := ""
	switch options.ValueStr("method") {
	case "":
		method = ""
	default:
		method = `<input type="hidden" name="_METHOD" value="` + options.ValueStr("method") + `"/>`
	}

	id := ""
	switch options.ValueStr("id") {
	case "":
		id = ""
	default:
		id = "/" + options.ValueStr("id")
	}

	resultBefore := `<form action="` + action + id + `" method="POST" `+ enctypeAtt +`>`
	frame := options.NewDataFrame()
	frame.Set("f", changeset)
	resultAfter := `</form>`
	result = resultBefore + method + options.FnData(frame) + resultAfter

	return
}
