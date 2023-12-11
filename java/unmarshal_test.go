package java

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	javalang "github.com/edutko/cafegopher/java/lang"
)

func TestUnmarshal(t *testing.T) {
	var name string
	var expected any

	name = "Boolean"
	expected = javalang.Boolean{Value: true}
	t.Run(name, func(t *testing.T) {
		var actual javalang.Boolean
		assert.Nil(t, Unmarshal(mustReadFile(name), &actual))
		assert.Equal(t, expected, actual)
	})

	name = "Byte"
	expected = javalang.Byte{Value: -1}
	t.Run(name, func(t *testing.T) {
		var actual javalang.Byte
		assert.Nil(t, Unmarshal(mustReadFile(name), &actual))
		assert.Equal(t, expected, actual)
	})

	expected = byte(255)
	t.Run("Byte as byte", func(t *testing.T) {
		var actual byte
		assert.Nil(t, Unmarshal(mustReadFile(name), &actual))
		assert.Equal(t, expected, actual)
	})

	name = "Character"
	expected = javalang.Character{Value: 'a'}
	t.Run(name, func(t *testing.T) {
		var actual javalang.Character
		assert.Nil(t, Unmarshal(mustReadFile(name), &actual))
		assert.Equal(t, expected, actual)
	})

	name = "Double"
	expected = javalang.Double{Value: 2.718}
	t.Run(name, func(t *testing.T) {
		var actual javalang.Double
		assert.Nil(t, Unmarshal(mustReadFile(name), &actual))
		assert.Equal(t, expected, actual)
	})

	name = "Float"
	expected = javalang.Float{Value: 3.14}
	t.Run(name, func(t *testing.T) {
		var actual javalang.Float
		assert.Nil(t, Unmarshal(mustReadFile(name), &actual))
		assert.Equal(t, expected, actual)
	})

	name = "Integer"
	expected = javalang.Integer{Value: 1234567890}
	t.Run(name, func(t *testing.T) {
		var actual javalang.Integer
		assert.Nil(t, Unmarshal(mustReadFile(name), &actual))
		assert.Equal(t, expected, actual)
	})

	name = "Long"
	expected = javalang.Long{Value: 9876543210}
	t.Run(name, func(t *testing.T) {
		var actual javalang.Long
		assert.Nil(t, Unmarshal(mustReadFile(name), &actual))
		assert.Equal(t, expected, actual)
	})

	name = "Short"
	expected = javalang.Short{Value: 32767}
	t.Run(name, func(t *testing.T) {
		var actual javalang.Short
		assert.Nil(t, Unmarshal(mustReadFile(name), &actual))
		assert.Equal(t, expected, actual)
	})

	name = "bytes"
	expected = [4]byte{0x11, 0x22, 0x33, 0x44}
	t.Run(name, func(t *testing.T) {
		var actual [4]byte
		assert.Nil(t, Unmarshal(mustReadFile(name), &actual))
		assert.Equal(t, expected, actual)
	})

	name = "int"
	expected = 456789
	t.Run(name, func(t *testing.T) {
		var actual int
		assert.Nil(t, Unmarshal(mustReadFile(name), &actual))
		assert.Equal(t, expected, actual)
	})

	name = "int"
	expected = uint(456789)
	t.Run(name, func(t *testing.T) {
		var actual uint
		assert.Nil(t, Unmarshal(mustReadFile(name), &actual))
		assert.Equal(t, expected, actual)
	})

	name = "long-string"
	expected = strings.Repeat("0123456789abcdef", 4096)
	t.Run(name, func(t *testing.T) {
		var actual string
		assert.Nil(t, Unmarshal(mustReadFile(name), &actual))
		assert.Equal(t, expected, actual)
	})

	name = "object"
	expected = foo{
		B:    0x7f,
		Bool: true,
		C:    'e',
		D:    3.14,
		F:    2.718,
		I:    7,
		L:    5000000000,
		S:    32767,
		O:    "hello",
		A:    []byte{0x11, 0x22, 0x33},
	}
	t.Run(name, func(t *testing.T) {
		var actual foo
		assert.Nil(t, Unmarshal(mustReadFile(name), &actual))
		assert.Equal(t, expected, actual)
	})

	name = "object"
	expected = altfoo{
		B: 0x7f,
		I: 7,
		L: 5000000000,
		S: 32767,
	}
	t.Run(name, func(t *testing.T) {
		var actual altfoo
		assert.Nil(t, Unmarshal(mustReadFile(name), &actual))
		assert.Equal(t, expected, actual)
	})

	name = "SealedObjectForKeyProtector"
	expected = sealedObject{
		SealAlg:          "PBEWithMD5AndTripleDES",
		ParamsAlg:        "PBEWithMD5AndTripleDES",
		EncryptedContent: unhex("dde157cfd1670f6e76efb93953cfdd4737bc24f9098253e8defb40a8e3e428efd025224f3ad8dba5e71bd16eccfa429a21fbcaa4cb8b5050866014fbaf4be4c3cbfe77aecf4438437b054da882b73e766020f83628e5d85d0fd03b5cc36d8ca8b294aa92afb8efdf3d55fb57683b0cd80c43308ce63552c49c05824980ae1e0122bddcb4f4016e9f324b88c535f90fa2ac2e53119a38721d6012af87e7f40522"),
		EncodedParams:    unhex("300f0408a47298a326aba6500203030d40"),
	}
	t.Run(name, func(t *testing.T) {
		var actual sealedObject
		assert.Nil(t, Unmarshal(mustReadFile(name), &actual))
		assert.Equal(t, expected, actual)
	})

	name = "string"
	expected = "hi"
	t.Run(name, func(t *testing.T) {
		var actual string
		assert.Nil(t, Unmarshal(mustReadFile(name), &actual))
		assert.Equal(t, expected, actual)
	})

	name = "strings"
	expected = []string{"abc", "def", "ghi"}
	t.Run(name, func(t *testing.T) {
		var actual []string
		assert.Nil(t, Unmarshal(mustReadFile(name), &actual))
		assert.Equal(t, expected, actual)
	})
}

func TestUnmarshal_errors(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		err := Unmarshal(mustReadFile("string"), nil)
		assert.NotNil(t, err)
	})

	t.Run("non-pointer", func(t *testing.T) {
		var actual string
		err := Unmarshal(mustReadFile("string"), actual)
		assert.NotNil(t, err)
	})

	t.Run("invalid data", func(t *testing.T) {
		var actual string
		err := Unmarshal([]byte{0x0b, 0xad, 0xf0, 0x0d}, &actual)
		assert.NotNil(t, err)
	})

	t.Run("type mismatch bool", func(t *testing.T) {
		var actual struct {
			Val bool `java:"com.edutko.Main$Foo/i"`
		}
		err := Unmarshal(mustReadFile("object"), &actual)
		assert.Equal(t, "unmarshalValue: cannot cast int to bool", err.Error())
	})

	t.Run("type mismatch int8", func(t *testing.T) {
		var actual struct {
			Val int8 `java:"com.edutko.Main$Foo/i"`
		}
		err := Unmarshal(mustReadFile("object"), &actual)
		assert.Equal(t, "unmarshalValue: cannot cast int to int8", err.Error())
	})

	t.Run("type mismatch int16", func(t *testing.T) {
		var actual struct {
			Val int16 `java:"com.edutko.Main$Foo/bool"`
		}
		err := Unmarshal(mustReadFile("object"), &actual)
		assert.Equal(t, "unmarshalValue: cannot cast bool to int16", err.Error())
	})

	t.Run("type mismatch int", func(t *testing.T) {
		var actual struct {
			Val int `java:"com.edutko.Main$Foo/bool"`
		}
		err := Unmarshal(mustReadFile("object"), &actual)
		assert.Equal(t, "unmarshalValue: cannot cast bool to int", err.Error())
	})

	t.Run("type mismatch rune", func(t *testing.T) {
		var actual struct {
			Val rune `java:"com.edutko.Main$Foo/bool"`
		}
		err := Unmarshal(mustReadFile("object"), &actual)
		assert.Equal(t, "unmarshalValue: cannot cast bool to int32", err.Error())
	})

	t.Run("type mismatch int64", func(t *testing.T) {
		var actual struct {
			Val int64 `java:"com.edutko.Main$Foo/bool"`
		}
		err := Unmarshal(mustReadFile("object"), &actual)
		assert.Equal(t, "unmarshalValue: cannot cast bool to int64", err.Error())
	})

	t.Run("type mismatch uint8", func(t *testing.T) {
		var actual struct {
			Val uint8 `java:"com.edutko.Main$Foo/i"`
		}
		err := Unmarshal(mustReadFile("object"), &actual)
		assert.Equal(t, "unmarshalValue: cannot cast int to uint8", err.Error())
	})

	t.Run("type mismatch uint16", func(t *testing.T) {
		var actual struct {
			Val uint16 `java:"com.edutko.Main$Foo/bool"`
		}
		err := Unmarshal(mustReadFile("object"), &actual)
		assert.Equal(t, "unmarshalValue: cannot cast bool to uint16", err.Error())
	})

	t.Run("type mismatch uint", func(t *testing.T) {
		var actual struct {
			Val uint `java:"com.edutko.Main$Foo/bool"`
		}
		err := Unmarshal(mustReadFile("object"), &actual)
		assert.Equal(t, "unmarshalValue: cannot cast bool to uint", err.Error())
	})

	t.Run("type mismatch uint32", func(t *testing.T) {
		var actual struct {
			Val uint32 `java:"com.edutko.Main$Foo/bool"`
		}
		err := Unmarshal(mustReadFile("object"), &actual)
		assert.Equal(t, "unmarshalValue: cannot cast bool to uint32", err.Error())
	})

	t.Run("type mismatch uint64", func(t *testing.T) {
		var actual struct {
			Val uint64 `java:"com.edutko.Main$Foo/bool"`
		}
		err := Unmarshal(mustReadFile("object"), &actual)
		assert.Equal(t, "unmarshalValue: cannot cast bool to uint64", err.Error())
	})

	t.Run("type mismatch float32", func(t *testing.T) {
		var actual struct {
			Val float32 `java:"com.edutko.Main$Foo/i"`
		}
		err := Unmarshal(mustReadFile("object"), &actual)
		assert.Equal(t, "unmarshalValue: cannot cast int to float32", err.Error())
	})

	t.Run("type mismatch float64", func(t *testing.T) {
		var actual struct {
			Val float64 `java:"com.edutko.Main$Foo/i"`
		}
		err := Unmarshal(mustReadFile("object"), &actual)
		assert.Equal(t, "unmarshalValue: cannot cast int to float64", err.Error())
	})

	t.Run("type mismatch array", func(t *testing.T) {
		var actual struct {
			Val [100]string `java:"com.edutko.Main$Foo/i"`
		}
		err := Unmarshal(mustReadFile("object"), &actual)
		assert.Equal(t, "unmarshalValue: cannot cast int to array", err.Error())
	})

	t.Run("type mismatch slice", func(t *testing.T) {
		var actual struct {
			Val []string `java:"com.edutko.Main$Foo/i"`
		}
		err := Unmarshal(mustReadFile("object"), &actual)
		assert.Equal(t, "unmarshalValue: cannot cast int to slice", err.Error())
	})

	t.Run("type mismatch string", func(t *testing.T) {
		var actual struct {
			Val string `java:"com.edutko.Main$Foo/i"`
		}
		err := Unmarshal(mustReadFile("object"), &actual)
		assert.Equal(t, "unmarshalValue: cannot cast int to string", err.Error())
	})

	t.Run("array size mismatch", func(t *testing.T) {
		var actual [2]string
		err := Unmarshal(mustReadFile("strings"), &actual)
		assert.Equal(t, "unmarshalValue: array size mismatch: expected 2, got 3", err.Error())
	})

	t.Run("unknown class", func(t *testing.T) {
		var actual struct {
			Val bool `java:"nonexistent"`
		}
		err := Unmarshal(mustReadFile("object"), &actual)
		assert.Equal(t, `unmarshalValue: getField: unable to determine class name for field "nonexistent"`, err.Error())
	})

	t.Run("nonexistent class", func(t *testing.T) {
		var actual struct {
			Val bool `java:"nonexistent/foo"`
		}
		err := Unmarshal(mustReadFile("object"), &actual)
		assert.Equal(t, `unmarshalValue: getField: "nonexistent": no such class in object`, err.Error())
	})

	t.Run("nonexistent field", func(t *testing.T) {
		var actual struct {
			Val bool `java:"com.edutko.Main$Foo/nonexistent"`
		}
		err := Unmarshal(mustReadFile("object"), &actual)
		assert.Equal(t, `unmarshalValue: getField: "nonexistent": no such field in object`, err.Error())
	})

	t.Run("untagged field", func(t *testing.T) {
		var actual struct {
			Val bool
		}
		err := Unmarshal(mustReadFile("object"), &actual)
		assert.Equal(t, `unmarshalValue: getField: unable to determine class name for field ""`, err.Error())
	})
}

func mustReadFile(name string) []byte {
	data, err := os.ReadFile(filepath.Join("testdata", name+".ser"))
	if err != nil {
		panic(err)
	}
	return data
}

type foo struct {
	B    byte    `java:"com.edutko.Main$Foo/b"`
	Bool bool    `java:"com.edutko.Main$Foo/bool"`
	C    rune    `java:"com.edutko.Main$Foo/c"`
	D    float64 `java:"com.edutko.Main$Foo/d"`
	F    float32 `java:"com.edutko.Main$Foo/f"`
	I    int     `java:"com.edutko.Main$Foo/i"`
	L    int64   `java:"com.edutko.Main$Foo/l"`
	S    int16   `java:"com.edutko.Main$Foo/s"`
	O    string  `java:"com.edutko.Main$Foo/o"`
	A    []byte  `java:"com.edutko.Main$Foo/a"`
}

type altfoo struct {
	B int8   `java:"com.edutko.Main$Foo/b"`
	I uint32 `java:"com.edutko.Main$Foo/i"`
	L uint64 `java:"com.edutko.Main$Foo/l"`
	S uint16 `java:"com.edutko.Main$Foo/s"`
}

type sealedObject struct {
	SealAlg          string `java:"javax.crypto.SealedObject/sealAlg"`
	ParamsAlg        string `java:"javax.crypto.SealedObject/paramsAlg"`
	EncryptedContent []byte `java:"javax.crypto.SealedObject/encryptedContent"`
	EncodedParams    []byte `java:"javax.crypto.SealedObject/encodedParams"`
}
