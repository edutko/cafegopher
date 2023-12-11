package java

import (
	"reflect"

	javalang "github.com/edutko/cafegopher/java/lang"
)

func AllowInt8ToByteCoercion() {
	allowInt8ToByteCoercion = true
}

func PreventInt8ToByteCoercion() {
	allowInt8ToByteCoercion = false
}

var allowInt8ToByteCoercion = true

func castToBool(x any) (bool, bool) {
	switch v := x.(type) {
	case bool:
		return v, true
	case Object:
		if val, ok := unmarshalWrappedPrimitive(v); ok {
			return castToBool(val)
		}
	}
	return false, false
}

func castToFloat32(x any) (float32, bool) {
	switch v := x.(type) {
	case float32:
		return v, true
	case Object:
		if val, ok := unmarshalWrappedPrimitive(v); ok {
			return castToFloat32(val)
		}
	}
	return 0, false
}

func castToFloat64(x any) (float64, bool) {
	switch v := x.(type) {
	case float64:
		return v, true
	case float32:
		return float64(v), true
	case Object:
		if val, ok := unmarshalWrappedPrimitive(v); ok {
			return castToFloat64(val)
		}
	}
	return 0, false
}

func castToInt(x any) (int, bool) {
	switch v := x.(type) {
	case int:
		return v, true
	case int32:
		return int(v), true
	case uint16:
		return int(v), true
	case int16:
		return int(v), true
	case uint8:
		return int(v), true
	case int8:
		return int(v), true
	case Object:
		if val, ok := unmarshalWrappedPrimitive(v); ok {
			return castToInt(val)
		}
	}
	return 0, false
}

func castToInt8(x any) (int8, bool) {
	switch v := x.(type) {
	case int8:
		return v, true
	case Object:
		if val, ok := unmarshalWrappedPrimitive(v); ok {
			return castToInt8(val)
		}
	}
	return 0, false
}

func castToInt16(x any) (int16, bool) {
	switch v := x.(type) {
	case int16:
		return v, true
	case uint8:
		return int16(v), true
	case int8:
		return int16(v), true
	case Object:
		if val, ok := unmarshalWrappedPrimitive(v); ok {
			return castToInt16(val)
		}
	}
	return 0, false
}

func castToInt32(x any) (int32, bool) {
	switch v := x.(type) {
	case int32:
		return v, true
	case int:
		return int32(v), true
	case uint16:
		return int32(v), true
	case int16:
		return int32(v), true
	case uint8:
		return int32(v), true
	case int8:
		return int32(v), true
	case Object:
		if val, ok := unmarshalWrappedPrimitive(v); ok {
			return castToInt32(val)
		}
	}
	return 0, false
}

func castToInt64(x any) (int64, bool) {
	switch v := x.(type) {
	case int64:
		return v, true
	case uint:
		return int64(v), true
	case int:
		return int64(v), true
	case uint32:
		return int64(v), true
	case int32:
		return int64(v), true
	case uint16:
		return int64(v), true
	case int16:
		return int64(v), true
	case uint8:
		return int64(v), true
	case int8:
		return int64(v), true
	case Object:
		if val, ok := unmarshalWrappedPrimitive(v); ok {
			return castToInt64(val)
		}
	}
	return 0, false
}

func castToUint(x any) (uint, bool) {
	switch v := x.(type) {
	case uint:
		return v, true
	case int:
		if v >= 0 {
			return uint(v), true
		}
	case uint32:
		return uint(v), true
	case int32:
		if v >= 0 {
			return uint(v), true
		}
	case uint16:
		return uint(v), true
	case int16:
		if v >= 0 {
			return uint(v), true
		}
	case uint8:
		return uint(v), true
	case int8:
		if v >= 0 {
			return uint(v), true
		}
	case Object:
		if val, ok := unmarshalWrappedPrimitive(v); ok {
			return castToUint(val)
		}
	}
	return 0, false
}

func castToUint8(x any) (uint8, bool) {
	switch v := x.(type) {
	case uint8:
		return v, true
	case int8:
		if allowInt8ToByteCoercion || v >= 0 {
			return uint8(v), true
		}
	case Object:
		if val, ok := unmarshalWrappedPrimitive(v); ok {
			return castToUint8(val)
		}
	}
	return 0, false
}

func castToUint16(x any) (uint16, bool) {
	switch v := x.(type) {
	case uint16:
		return v, true
	case int16:
		if v >= 0 {
			return uint16(v), true
		}
	case uint8:
		return uint16(v), true
	case int8:
		if v >= 0 {
			return uint16(v), true
		}
	case Object:
		if val, ok := unmarshalWrappedPrimitive(v); ok {
			return castToUint16(val)
		}
	}
	return 0, false
}

func castToUint32(x any) (uint32, bool) {
	switch v := x.(type) {
	case uint32:
		return v, true
	case uint:
		return uint32(v), true
	case int:
		if v >= 0 {
			return uint32(v), true
		}
	case int32:
		if v >= 0 {
			return uint32(v), true
		}
	case uint16:
		return uint32(v), true
	case int16:
		if v >= 0 {
			return uint32(v), true
		}
	case uint8:
		return uint32(v), true
	case int8:
		if v >= 0 {
			return uint32(v), true
		}
	case Object:
		if val, ok := unmarshalWrappedPrimitive(v); ok {
			return castToUint32(val)
		}
	}
	return 0, false
}

func castToUint64(x any) (uint64, bool) {
	switch v := x.(type) {
	case uint64:
		return v, true
	case int64:
		if v >= 0 {
			return uint64(v), true
		}
	case uint:
		return uint64(v), true
	case int:
		if v >= 0 {
			return uint64(v), true
		}
	case uint32:
		return uint64(v), true
	case int32:
		if v >= 0 {
			return uint64(v), true
		}
	case uint16:
		return uint64(v), true
	case int16:
		if v >= 0 {
			return uint64(v), true
		}
	case uint8:
		return uint64(v), true
	case int8:
		if v >= 0 {
			return uint64(v), true
		}
	case Object:
		if val, ok := unmarshalWrappedPrimitive(v); ok {
			return castToUint64(val)
		}
	}
	return 0, false
}

func unmarshalWrappedPrimitive(object Object) (any, bool) {
	switch object.ClassDesc.ClassName {
	case "java.lang.Boolean":
		var v javalang.Boolean
		if err := unmarshalValue(object, reflect.ValueOf(&v)); err == nil {
			return v.Value, true
		}
	case "java.lang.Byte":
		var v javalang.Byte
		if err := unmarshalValue(object, reflect.ValueOf(&v)); err == nil {
			return v.Value, true
		}
	case "java.lang.Double":
		var v javalang.Double
		if err := unmarshalValue(object, reflect.ValueOf(&v)); err == nil {
			return v.Value, true
		}
	case "java.lang.Float":
		var v javalang.Float
		if err := unmarshalValue(object, reflect.ValueOf(&v)); err == nil {
			return v.Value, true
		}
	case "java.lang.Integer":
		var v javalang.Integer
		if err := unmarshalValue(object, reflect.ValueOf(&v)); err == nil {
			return v.Value, true
		}
	case "java.lang.Long":
		var v javalang.Long
		if err := unmarshalValue(object, reflect.ValueOf(&v)); err == nil {
			return v.Value, true
		}
	case "java.lang.Short":
		var v javalang.Short
		if err := unmarshalValue(object, reflect.ValueOf(&v)); err == nil {
			return v.Value, true
		}
	}

	return nil, false
}
