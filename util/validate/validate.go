package validate

import (
	"reflect"
	"time"
)

func ValuesHasZero(vs ...interface{}) bool {
	for _, v := range vs {
		obj := reflect.ValueOf(v)
		if obj.IsZero() {
			return true
		}
	}

	return false
}

// StringToDate string_format: 2006-01-02 15:04:05, if it has error, this will return zero
func StringToDate(src string) time.Time {
	t, ok := time.Parse("2006-01-02 15:04:05", src)
	if ok != nil {
		return time.Time{}
	}

	return t
}