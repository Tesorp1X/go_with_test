package reflection

import "reflect"

func walk(x any, fn func(string)) {
	val := getValue(x)

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := range val.NumField() {
			walk(val.Field(i).Interface(), fn)
		}
	case reflect.Slice, reflect.Array:
		for i := range val.Len() {
			walk(val.Index(i).Interface(), fn)
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walk(val.MapIndex(key).Interface(), fn)
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
