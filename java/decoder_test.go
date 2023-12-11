package java

import (
	"encoding/hex"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeserializer_ReadContents(t *testing.T) {
	testCases := []struct {
		name     string
		expected []Content
	}{
		{"ArrayList", []Content{al}},
		{"bytes", []Content{Array{ClassDesc: &byteArray, Values: byteArrayValue("11223344")}}},
		{"enum", []Content{Enum{ClassDesc: &comEdutkoMainStatus, ConstantName: "FUBAR"}}},
		{"int", []Content{Object{
			ClassDesc: &javaLangInteger,
			ClassData: map[string]ClassData{"java.lang.Number": {}, "java.lang.Integer": {"value": 456789}},
		}}},
		{"object", []Content{foo1}},
		{"objects", []Content{Array{ClassDesc: &comEdutkoMainFooArray, Values: []Value{foo1, foo2, foo3}}}},
		{"SealedObjectForKeyProtector", []Content{secretKey}},
		{"long-string", []Content{strings.Repeat("0123456789abcdef", 4096)}},
		{"string", []Content{"hi"}},
		{"strings", []Content{Array{ClassDesc: &stringArray, Values: []Value{"abc", "def", "ghi"}}}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			f, err := os.Open(filepath.Join("testdata", tc.name+".ser"))
			if err != nil {
				panic(err)
			}
			d := NewDecoder(f)

			actual, err := d.DecodeAll(1)

			assert.Nil(t, err)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func byteArrayValue(s string) []Value {
	bs, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	var vs []Value
	for _, b := range bs {
		vs = append(vs, Value(int8(b)))
	}
	return vs
}

var comEdutkoMainFoo = Class{
	ClassName:        "com.edutko.Main$Foo",
	SerialVersionUID: 4038854518001758741,
	Info: ClassDescInfo{
		Flags: 0x02,
		Fields: []Field{
			{TypeByte, "b", ""},
			{TypeBoolean, "bool", ""},
			{TypeChar, "c", ""},
			{TypeDouble, "d", ""},
			{TypeFloat, "f", ""},
			{TypeInteger, "i", ""},
			{TypeLong, "l", ""},
			{TypeShort, "s", ""},
			{TypeArray, "a", "[B"},
			{TypeArray, "bars", "[Lcom/edutko/Main$Bar;"},
			{TypeObject, "o", "Ljava/lang/String;"},
			{TypeObject, "prefix", "Lcom/edutko/Main$Prefix;"},
			{TypeObject, "status", "Lcom/edutko/Main$Status;"},
		},
	},
}

var comEdutkoMainBar = Class{
	ClassName:        "com.edutko.Main$Bar",
	SerialVersionUID: 5937967780743061417,
	Info: ClassDescInfo{
		Flags: 0x02,
		Fields: []Field{
			{TypeInteger, "value", ""},
		},
	},
}

var comEdutkoMainBarArray = Class{
	ClassName:        "[Lcom.edutko.Main$Bar;",
	SerialVersionUID: 5949689868044538634,
	Info:             ClassDescInfo{Flags: 0x02},
}

var comEdutkoMainFooArray = Class{
	ClassName:        "[Lcom.edutko.Main$Foo;",
	SerialVersionUID: -3106264638368614811,
	Info:             ClassDescInfo{Flags: 0x02},
}

var comEdutkoMainPrefix = Class{
	ClassName: "com.edutko.Main$Prefix",
	Info: ClassDescInfo{
		Flags:          0x12,
		SuperClassDesc: &javaLangEnum,
	},
}

var comEdutkoMainStatus = Class{
	ClassName: "com.edutko.Main$Status",
	Info: ClassDescInfo{
		Flags:          0x12,
		SuperClassDesc: &javaLangEnum,
	},
}

var al = Object{
	ClassDesc: &javaUtilArrayList,
	ClassData: map[string]ClassData{
		"java.util.ArrayList": {
			"size": 4,
			"[object annotation]": []Annotation{
				blockData{0x00, 0x00, 0x00, 0x04},
				javaInteger(0xaaaaaa),
				javaInteger(0xbbbbbb),
				javaInteger(0xcccccc),
				javaInteger(0xdddddd),
			},
		},
	},
}

var foo1 = Object{
	ClassDesc: &comEdutkoMainFoo,
	ClassData: map[string]ClassData{
		"com.edutko.Main$Foo": {
			"b":    int8(0x7f),
			"bool": true,
			"c":    'e',
			"d":    3.14,
			"f":    float32(2.718),
			"i":    7,
			"l":    int64(5000000000),
			"s":    int16(32767),
			"o":    "hello",
			"a": Array{
				ClassDesc: &byteArray,
				Values:    byteArrayValue("112233"),
			},
			"prefix": Enum{
				ClassDesc:    &comEdutkoMainPrefix,
				ConstantName: "KILO",
			},
			"status": Enum{
				ClassDesc:    &comEdutkoMainStatus,
				ConstantName: "SNAFU",
			},
			"bars": Array{
				ClassDesc: &comEdutkoMainBarArray,
				Values: []Value{
					Object{
						ClassDesc: &comEdutkoMainBar,
						ClassData: map[string]ClassData{
							"com.edutko.Main$Bar": {"value": 0x1111},
						},
					},
					Object{
						ClassDesc: &comEdutkoMainBar,
						ClassData: map[string]ClassData{
							"com.edutko.Main$Bar": {"value": 0x2222},
						},
					},
				},
			},
		},
	},
}

var foo2 = Object{
	ClassDesc: &comEdutkoMainFoo,
	ClassData: map[string]ClassData{
		"com.edutko.Main$Foo": {
			"b":    int8(0x00),
			"bool": true,
			"c":    'e',
			"d":    3.14,
			"f":    float32(2.718),
			"i":    8,
			"l":    int64(6000000000),
			"s":    int16(32766),
			"o":    "hola",
			"a": Array{
				ClassDesc: &byteArray,
				Values:    byteArrayValue("445566"),
			},
			"prefix": Enum{
				ClassDesc:    &comEdutkoMainPrefix,
				ConstantName: "MEGA",
			},
			"status": Enum{
				ClassDesc:    &comEdutkoMainStatus,
				ConstantName: "TARFU",
			},
			"bars": Array{
				ClassDesc: &comEdutkoMainBarArray,
				Values: []Value{
					Object{
						ClassDesc: &comEdutkoMainBar,
						ClassData: map[string]ClassData{
							"com.edutko.Main$Bar": {"value": 0x3333},
						},
					},
					Object{
						ClassDesc: &comEdutkoMainBar,
						ClassData: map[string]ClassData{
							"com.edutko.Main$Bar": {"value": 0x4444},
						},
					},
				},
			},
		},
	},
}

var foo3 = Object{
	ClassDesc: &comEdutkoMainFoo,
	ClassData: map[string]ClassData{
		"com.edutko.Main$Foo": {
			"b":    int8(0x55),
			"bool": true,
			"c":    'e',
			"d":    3.14,
			"f":    float32(2.718),
			"i":    9,
			"l":    int64(7000000000),
			"s":    int16(32765),
			"o":    "aloha",
			"a": Array{
				ClassDesc: &byteArray,
				Values:    byteArrayValue("777777"),
			},
			"prefix": Enum{
				ClassDesc:    &comEdutkoMainPrefix,
				ConstantName: "GIGA",
			},
			"status": Enum{
				ClassDesc:    &comEdutkoMainStatus,
				ConstantName: "FUBAR",
			},
			"bars": Array{
				ClassDesc: &comEdutkoMainBarArray,
				Values: []Value{
					Object{
						ClassDesc: &comEdutkoMainBar,
						ClassData: map[string]ClassData{
							"com.edutko.Main$Bar": {"value": 0x5555},
						},
					},
				},
			},
		},
	},
}

var secretKey = Object{
	ClassDesc: &comSunCryptoProviderSealedObjectForKeyProtector,
	ClassData: map[string]ClassData{
		"javax.crypto.SealedObject": {
			"encodedParams": Array{
				ClassDesc: &byteArray,
				Values:    byteArrayValue("300f0408a47298a326aba6500203030d40"),
			},
			"encryptedContent": Array{
				ClassDesc: &byteArray,
				Values:    byteArrayValue("dde157cfd1670f6e76efb93953cfdd4737bc24f9098253e8defb40a8e3e428efd025224f3ad8dba5e71bd16eccfa429a21fbcaa4cb8b5050866014fbaf4be4c3cbfe77aecf4438437b054da882b73e766020f83628e5d85d0fd03b5cc36d8ca8b294aa92afb8efdf3d55fb57683b0cd80c43308ce63552c49c05824980ae1e0122bddcb4f4016e9f324b88c535f90fa2ac2e53119a38721d6012af87e7f40522"),
			},
			"paramsAlg": "PBEWithMD5AndTripleDES",
			"sealAlg":   "PBEWithMD5AndTripleDES",
		},
		"com.sun.crypto.provider.SealedObjectForKeyProtector": {}},
}
