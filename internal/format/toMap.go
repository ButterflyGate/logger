package format

// import (
// 	"reflect"
// 	"strings"
// )

// func structToMap(data interface{}) interface{} {
// 	if data == nil {
// 		return data
// 	}
// 	ret := make(map[string]interface{}, 30)
// 	val := reflect.ValueOf(data)
// 	for i := 0; i < val.NumField(); i++ {
// 		v := val.Field(i)
// 		e := v.Elem()
// 		k := v.
// 		if tag := f.Tag.Get("json"); tag != "" {
// 			if strings.Contains(tag, `"omitempty"`) && isZeroValue[reflect.Value](val) {

// 			}
// 			k = tag
// 		}
// 	}
// 	return ret
// }

// func isZeroValue[T comparable](val T) bool {

// 	var zero T
// 	return zero == val
// }
