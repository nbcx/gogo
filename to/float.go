package to

import (
	"reflect"
)

var floatType = reflect.TypeOf(float64(0))

// Float64 guess Num format and convert to Float64
func Float64(unk interface{}) float64 {
	v := reflect.ValueOf(unk)
	v = reflect.Indirect(v)
	if !v.Type().ConvertibleTo(floatType) {
		return 0
	}
	fv := v.Convert(floatType)
	return fv.Float()
}
