package java

import "encoding/hex"

func javaBoolean(val bool) Object {
	return Object{
		ClassDesc: &javaLangBoolean,
		ClassData: ClassData{
			javaLangBoolean.ClassName: {"value": val},
		},
	}
}

func javaByte(val int8) Object {
	return Object{
		ClassDesc: &javaLangByte,
		ClassData: ClassData{
			javaLangNumber.ClassName: {},
			javaLangByte.ClassName:   {"value": val},
		},
	}
}

func javaCharacter(val rune) Object {
	return Object{
		ClassDesc: &javaLangCharacter,
		ClassData: ClassData{
			javaLangCharacter.ClassName: {"value": val},
		},
	}
}

func javaDouble(val float64) Object {
	return Object{
		ClassDesc: &javaLangDouble,
		ClassData: ClassData{
			javaLangNumber.ClassName: {},
			javaLangDouble.ClassName: {"value": val},
		},
	}
}

func javaFloat(val float32) Object {
	return Object{
		ClassDesc: &javaLangFloat,
		ClassData: ClassData{
			javaLangNumber.ClassName: {},
			javaLangFloat.ClassName:  {"value": val},
		},
	}
}

func javaInteger(val int) Object {
	return Object{
		ClassDesc: &javaLangInteger,
		ClassData: ClassData{
			javaLangNumber.ClassName:  {},
			javaLangInteger.ClassName: {"value": val},
		},
	}
}

func javaLong(val int64) Object {
	return Object{
		ClassDesc: &javaLangLong,
		ClassData: ClassData{
			javaLangNumber.ClassName: {},
			javaLangLong.ClassName:   {"value": val},
		},
	}
}

func javaShort(val int16) Object {
	return Object{
		ClassDesc: &javaLangShort,
		ClassData: ClassData{
			javaLangNumber.ClassName: {},
			javaLangShort.ClassName:  {"value": val},
		},
	}
}

func unhex(s string) []byte {
	if b, err := hex.DecodeString(s); err != nil {
		panic(err)
	} else {
		return b
	}
}

var byteArray = Class{
	ClassName:        "[B",
	SerialVersionUID: -5984413125824719648,
	Info:             ClassDescInfo{Flags: 0x02},
}

var stringArray = Class{
	ClassName:        "[Ljava.lang.String;",
	SerialVersionUID: -5921575005990323385,
	Info:             ClassDescInfo{Flags: 0x02},
}

var javaLangBoolean = Class{
	ClassName:        "java.lang.Boolean",
	SerialVersionUID: -3665804199014368530,
	Info: ClassDescInfo{
		Flags: 0x02,
		Fields: []Field{
			{TypeCode: TypeBoolean, FieldName: "value"},
		},
	},
}

var javaLangByte = Class{
	ClassName:        "java.lang.Byte",
	SerialVersionUID: -7183698231559129828,
	Info: ClassDescInfo{
		Flags: 0x02,
		Fields: []Field{
			{TypeCode: TypeByte, FieldName: "value"},
		},
	},
}

var javaLangCharacter = Class{
	ClassName:        "java.lang.Character",
	SerialVersionUID: 3786198910865385080,
	Info: ClassDescInfo{
		Flags: 0x02,
		Fields: []Field{
			{TypeCode: TypeChar, FieldName: "value"},
		},
	},
}

var javaLangDouble = Class{
	ClassName:        "java.lang.Double",
	SerialVersionUID: -9172774392245257468,
	Info: ClassDescInfo{
		Flags: 0x02,
		Fields: []Field{
			{TypeCode: TypeDouble, FieldName: "value"},
		},
	},
}

var javaLangFloat = Class{
	ClassName:        "java.lang.Float",
	SerialVersionUID: -2671257302660747028,
	Info: ClassDescInfo{
		Flags: 0x02,
		Fields: []Field{
			{TypeCode: TypeFloat, FieldName: "value"},
		},
	},
}

var javaLangEnum = Class{
	ClassName: "java.lang.Enum",
	Info:      ClassDescInfo{Flags: 0x12},
}

var javaLangInteger = Class{
	ClassName:        "java.lang.Integer",
	SerialVersionUID: 1360826667806852920,
	Info: ClassDescInfo{
		Flags: 0x02,
		Fields: []Field{
			{TypeCode: TypeInteger, FieldName: "value"},
		},
		SuperClassDesc: &javaLangNumber,
	},
}

var javaLangLong = Class{
	ClassName:        "java.lang.Long",
	SerialVersionUID: 4290774380558885855,
	Info: ClassDescInfo{
		Flags: 0x02,
		Fields: []Field{
			{TypeCode: TypeByte, FieldName: "value"},
		},
	},
}

var javaLangNumber = Class{
	ClassName:        "java.lang.Number",
	SerialVersionUID: -8742448824652078965,
	Info:             ClassDescInfo{Flags: 0x02},
}

var javaLangShort = Class{
	ClassName:        "java.lang.Short",
	SerialVersionUID: 7515723908773894738,
	Info: ClassDescInfo{
		Flags: 0x02,
		Fields: []Field{
			{TypeCode: TypeByte, FieldName: "value"},
		},
	},
}

var javaUtilArrayList = Class{
	ClassName:        "java.util.ArrayList",
	SerialVersionUID: 8683452581122892189,
	Info: ClassDescInfo{
		Flags:  0x03,
		Fields: []Field{{TypeCode: TypeInteger, FieldName: "size"}},
	},
}

var comSunCryptoProviderSealedObjectForKeyProtector = Class{
	ClassName:        "com.sun.crypto.provider.SealedObjectForKeyProtector",
	SerialVersionUID: -3650226485480866989,
	Info: ClassDescInfo{
		Flags:          0x02,
		SuperClassDesc: &javaxCryptoSealedObject,
	},
}

var javaxCryptoSealedObject = Class{
	ClassName:        "javax.crypto.SealedObject",
	SerialVersionUID: 4482838265551344752,
	Info: ClassDescInfo{
		Flags: 0x02,
		Fields: []Field{
			{TypeCode: TypeArray, FieldName: "encodedParams", ClassName: "[B"},
			{TypeCode: TypeArray, FieldName: "encryptedContent", ClassName: "[B"},
			{TypeCode: TypeObject, FieldName: "paramsAlg", ClassName: "Ljava/lang/String;"},
			{TypeCode: TypeObject, FieldName: "sealAlg", ClassName: "Ljava/lang/String;"},
		},
	},
}
