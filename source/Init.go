package goform

import "fmt"
import "syscall/js"

func Init(schemas *map[string]any) {

	forms := js.Global().Get("document").Get("querySelectorAll").Call("form[enctype=\"application/json\"")

	for f := 0; f < len(forms); f++ {

		form := forms[f]

		form_action := form.Get("getAttribute").Call("action").String()
		form_enctype := form.Get("getAttribute").Call("enctype").String()
		form_method := form.Get("getAttribute").Call("method").String()
		form_schema := form.Get("getAttribute").Call("data-schema").String()

		fmt.Println(form_action, form_enctype, form_method, form_schema)

		// TODO: How to get json tags from struct?
		// TODO: input data-name=json?

	}

}
