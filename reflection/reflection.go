package reflection

import "reflect"

func walk(x any, fn func(string)) {
	val := getValue(x)

	for i := range val.NumField() {
		field := val.Field(i)

		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			walk(field.Interface(), fn)
		}
	}
}

func getValue(x any) (val reflect.Value) {
	val = reflect.ValueOf(x)
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	return
}
