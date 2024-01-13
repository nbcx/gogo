package to

import (
	"reflect"
)

func Int64(v any) int64 {
	switch t := v.(type) {
	case int:
		return int64(t)
	case float64, float32:
		return int64(reflect.ValueOf(t).Float())
	case int64, int32:
		return reflect.ValueOf(t).Int()
	default:
		return 0
	}
}

func Int(v any) int {
	return int(Int64(v))
}
