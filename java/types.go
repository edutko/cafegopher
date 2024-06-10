package java

import (
	"fmt"
	"strings"
)

type Annotation any

type Array struct {
	ClassDesc *Class
	Values    []Value
}

func (a Array) Length() int {
	return len(a.Values)
}

func (a Array) Get(index int) Value {
	return a.Values[index]
}

func (a Array) ItemType() (TypeCode, string) {
	className := a.ClassDesc.ClassName
	if len(className) < 2 {
		return 0, ""
	}
	return TypeCode(className[1]), strings.TrimSuffix(strings.TrimPrefix(className, "["), ";")
}

type Class struct {
	ClassName        string
	SerialVersionUID SerialVersionUID
	Info             ClassDescInfo
}

type ClassData map[string]map[string]Value
type ClassDescFlags byte

const (
	ScWriteMethod    = 0x01
	ScBlockData      = 0x08
	ScSerializable   = 0x02
	ScExternalizable = 0x04
	ScEnum           = 0x10
)

func (f ClassDescFlags) IsSerializable() bool {
	return f&ScSerializable == ScSerializable
}

func (f ClassDescFlags) HasWriteMethod() bool {
	return f&ScWriteMethod == ScWriteMethod
}

type ClassDescInfo struct {
	Flags           ClassDescFlags
	Fields          []Field
	ClassAnnotation []Annotation
	SuperClassDesc  *Class
}

type Content any

type Enum struct {
	ClassDesc    *Class
	ConstantName string
}

type Field struct {
	TypeCode  TypeCode
	FieldName string
	ClassName string
}

type Object struct {
	ClassDesc *Class
	ClassData ClassData
}

func (o Object) GetClassName() string {
	if o.ClassDesc == nil {
		return ""
	}
	return o.ClassDesc.ClassName
}

func (o Object) GetField(name string) (any, error) {
	parts := strings.Split(name, ".")
	last := len(parts) - 1
	className := strings.Join(parts[0:last], ".")
	fieldName := parts[last]
	if fields, ok := o.ClassData[className]; ok {
		if value, ok := fields[fieldName]; ok {
			return value, nil
		}
		return nil, fmt.Errorf("%#v: %w", fieldName, ErrNoSuchField)
	}
	return nil, fmt.Errorf("%#v: %w", className, ErrNoSuchClass)
}

type SerialVersionUID int64

type TypeCode byte

func (t TypeCode) IsObject() bool {
	return t == TypeObject || t == TypeArray
}

func (t TypeCode) IsPrimitive() bool {
	return !t.IsObject()
}

const (
	TypeByte    TypeCode = 'B'
	TypeChar    TypeCode = 'C'
	TypeDouble  TypeCode = 'D'
	TypeFloat   TypeCode = 'F'
	TypeInteger TypeCode = 'I'
	TypeLong    TypeCode = 'J'
	TypeShort   TypeCode = 'S'
	TypeBoolean TypeCode = 'Z'
	TypeArray   TypeCode = '['
	TypeObject  TypeCode = 'L'
)

type Value any
