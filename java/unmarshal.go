package java

import (
	"bytes"
	"fmt"
	"io"
	"path"
	"reflect"
	"strings"
)

func AddPackagePrefixes(prefixes ...string) {
	packagePrefixes = append(packagePrefixes, prefixes...)
}

func SetPackagePrefixes(prefixes ...string) {
	packagePrefixes = prefixes
}

func Unmarshal(data []byte, v any) error {
	return UnmarshalReader(bytes.NewReader(data), v)
}

func UnmarshalReader(r io.Reader, v any) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Pointer || rv.IsNil() {
		return &InvalidUnmarshalError{reflect.TypeOf(v)}
	}

	content, err := NewDecoder(r).Decode()
	if err != nil {
		return fmt.Errorf("java.Decode: %w", err)
	}

	err = unmarshalValue(content, rv)
	if err != nil {
		return fmt.Errorf("unmarshalValue: %w", err)
	}

	return nil
}

func unmarshalValue(javaValue Content, goValue reflect.Value) error {
	switch k := goValue.Kind(); k {
	case reflect.Bool:
		if v, ok := castToBool(javaValue); ok {
			goValue.Set(reflect.ValueOf(v))
		} else {
			return fmt.Errorf("cannot cast %T to bool", javaValue)
		}

	case reflect.Int:
		if v, ok := castToInt(javaValue); ok {
			goValue.Set(reflect.ValueOf(v))
		} else {
			return fmt.Errorf("cannot cast %T to int", javaValue)
		}

	case reflect.Int8:
		if v, ok := castToInt8(javaValue); ok {
			goValue.Set(reflect.ValueOf(v))
		} else {
			return fmt.Errorf("cannot cast %T to int8", javaValue)
		}

	case reflect.Int16:
		if v, ok := castToInt16(javaValue); ok {
			goValue.Set(reflect.ValueOf(v))
		} else {
			return fmt.Errorf("cannot cast %T to int16", javaValue)
		}

	case reflect.Int32:
		if v, ok := castToInt32(javaValue); ok {
			goValue.Set(reflect.ValueOf(v))
		} else {
			return fmt.Errorf("cannot cast %T to int32", javaValue)
		}

	case reflect.Int64:
		if v, ok := castToInt64(javaValue); ok {
			goValue.Set(reflect.ValueOf(v))
		} else {
			return fmt.Errorf("cannot cast %T to int64", javaValue)
		}

	case reflect.Uint:
		if v, ok := castToUint(javaValue); ok {
			goValue.Set(reflect.ValueOf(v))
		} else {
			return fmt.Errorf("cannot cast %T to uint", javaValue)
		}

	case reflect.Uint8:
		if v, ok := castToUint8(javaValue); ok {
			goValue.Set(reflect.ValueOf(v))
		} else {
			return fmt.Errorf("cannot cast %T to uint8", javaValue)
		}

	case reflect.Uint16:
		if v, ok := castToUint16(javaValue); ok {
			goValue.Set(reflect.ValueOf(v))
		} else {

			return fmt.Errorf("cannot cast %T to uint16", javaValue)
		}

	case reflect.Uint32:
		if v, ok := castToUint32(javaValue); ok {
			goValue.Set(reflect.ValueOf(v))
		} else {
			return fmt.Errorf("cannot cast %T to uint32", javaValue)
		}

	case reflect.Uint64:
		if v, ok := castToUint64(javaValue); ok {
			goValue.Set(reflect.ValueOf(v))
		} else {
			return fmt.Errorf("cannot cast %T to uint64", javaValue)
		}

	case reflect.Float32:
		if v, ok := castToFloat32(javaValue); ok {
			goValue.Set(reflect.ValueOf(v))
		} else {
			return fmt.Errorf("cannot cast %T to float32", javaValue)
		}

	case reflect.Float64:
		if v, ok := castToFloat64(javaValue); ok {
			goValue.Set(reflect.ValueOf(v))
		} else {
			return fmt.Errorf("cannot cast %T to float64", javaValue)
		}

	case reflect.Array:
		if arr, ok := javaValue.(Array); ok {
			if len(arr.Values) != goValue.Len() {
				return fmt.Errorf("array size mismatch: expected %d, got %d", goValue.Len(), len(arr.Values))
			}
			for i := 0; i < arr.Length(); i++ {
				err := unmarshalValue(arr.Get(i), goValue.Index(i))
				if err != nil {
					return err
				}
			}
		} else {
			return fmt.Errorf("cannot cast %T to array", javaValue)
		}

	case reflect.Map:
		// TODO: implement support for Java HashMap
		// TODO: consider support for unmarshalling arbitrary objects as map[string]any
		return ErrNotSupported

	case reflect.Pointer:
		return unmarshalValue(javaValue, goValue.Elem())

	case reflect.Slice:
		if arr, ok := javaValue.(Array); ok {
			for i := 0; i < arr.Length(); i++ {
				itm := reflect.New(goValue.Type().Elem())
				err := unmarshalValue(arr.Get(i), itm)
				if err != nil {
					return err
				}
				goValue.Set(reflect.Append(goValue, reflect.Indirect(itm)))
			}
		} else {
			return fmt.Errorf("cannot cast %T to slice", javaValue)
		}

	case reflect.String:
		if v, ok := javaValue.(string); ok {
			goValue.Set(reflect.ValueOf(v))
		} else {
			return fmt.Errorf("cannot cast %T to string", javaValue)
		}

	case reflect.Struct:
		if javaObj, ok := javaValue.(Object); ok {
			typ := goValue.Type()
			for i := 0; i < typ.NumField(); i++ {
				sf := typ.Field(i)
				goField := goValue.FieldByName(sf.Name)
				if !goField.CanSet() {
					continue
				}
				javaField, err := getField(javaObj, typ, i)
				if err != nil {
					return fmt.Errorf("getField: %w", err)
				}
				err = unmarshalValue(javaField, goField)
				if err != nil {
					return err
				}
			}
		}

	case reflect.Chan, reflect.Complex64, reflect.Complex128, reflect.Func, reflect.Interface, reflect.Invalid, reflect.Uintptr, reflect.UnsafePointer:
		return ErrNotSupported

	default:
		panic(fmt.Sprintf("unknown kind: %d", k))
	}

	return nil
}

func getField(object Object, goType reflect.Type, index int) (Content, error) {
	goField := goType.Field(index)
	className, fieldName := splitTag(goField.Tag.Get("java"))
	if className == "" {
		pkgPath := goType.PkgPath()
		for _, p := range packagePrefixes {
			if strings.HasPrefix(pkgPath, p) {
				pkgPath = strings.TrimPrefix(pkgPath, p)
				break
			}
		}
		pkgPath = strings.TrimPrefix(pkgPath, "/")
		className = path.Join(pkgPath, goType.Name())
		className = strings.ReplaceAll(className, "/", ".")
	}

	if className == "" {
		return nil, fmt.Errorf("unable to determine class name for field %#v", goField.Tag.Get("java"))
	}
	values, exists := object.ClassData[className]
	if !exists {
		return nil, fmt.Errorf("%#v: %w", className, ErrNoSuchClass)
	}
	value, exists := values[fieldName]
	if !exists {
		return nil, fmt.Errorf("%#v: %w", fieldName, ErrNoSuchField)
	}

	return value, nil
}

func splitTag(tag string) (class, field string) {
	parts := strings.SplitN(tag, "/", 2)
	if len(parts) == 1 {
		return "", parts[0]
	}
	return parts[0], parts[1]
}

var packagePrefixes = []string{"github.com/edutko/cafegopher"}
