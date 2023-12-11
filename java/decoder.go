package java

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"math"
)

const (
	StreamMagic   = "\xAC\xED"
	StreamVersion = int16(5)
)

func NewDecoder(r io.Reader) *Decoder {
	d := &Decoder{r: &binaryReader{r}}
	d.Reset()
	return d
}

// Decode parses the first element of a stream of serialized Java objects
// https://docs.oracle.com/javase/6/docs/platform/serialization/spec/protocol.html
func (d *Decoder) Decode() (Content, error) {
	c, err := d.DecodeAll(1)
	if err != nil {
		return nil, err
	}
	if len(c) > 0 {
		return c[0], nil
	}
	return nil, nil
}

// DecodeAll parses serialized Java objects
// https://docs.oracle.com/javase/6/docs/platform/serialization/spec/protocol.html
func (d *Decoder) DecodeAll(limit int) ([]Content, error) {
	magic, err := d.r.readBytes(2)
	if err != nil {
		return nil, fmt.Errorf("readBytes: %w", err)
	}
	if !bytes.Equal(magic, []byte(StreamMagic)) {
		return nil, fmt.Errorf("invalid stream: incorrect magic")
	}

	version, err := d.r.readInt16()
	if err != nil {
		return nil, fmt.Errorf("readInt16: %w", err)
	}
	if version != StreamVersion {
		return nil, fmt.Errorf("unsupported serialization version: %d", version)
	}

	return d.readContents(limit)
}

func (d *Decoder) Reset() {
	d.h = baseHandleValue
	d.o = make(map[handle]Content)
}

type Decoder struct {
	r *binaryReader
	h handle
	o map[handle]Content
}

func (d *Decoder) readContents(limit int) ([]Content, error) {
	// contents:
	//   content
	//   contents content
	if limit < 0 {
		limit = math.MaxInt
	}
	var contents []Content
	for i := 0; i < limit; i++ {
		c, err := d.readContent()
		if err != nil {
			return contents, fmt.Errorf("d.readContent: %w", err)
		}
		if c != nil {
			contents = append(contents, c)
		}
	}
	return contents, nil
}

func (d *Decoder) readContent() (Content, error) {
	// content:
	//   object
	//   blockdata
	// object:
	//   newObject
	//   newClass
	//   newArray
	//   newString
	//   newEnum
	//   newClassDesc
	//   prevObject
	//   nullReference
	//   exception
	//   TC_RESET
	// nullReference:
	//   TC_NULL
	t, err := d.r.readByte()
	if err != nil {
		return nil, fmt.Errorf("readByte: %w", err)
	}

	switch t {
	case tcObject:
		return d.readNewObject()
	case tcClass:
		return d.readNewClass()
	case tcArray:
		return d.readNewArray()
	case tcString:
		return d.readNewString()
	case tcLongString:
		return d.readNewLongString()
	case tcEnum:
		return d.readNewEnum()
	case tcClassDesc:
		return d.readNewClassDesc()
	case tcProxyClassDesc:
		return d.readNewProxyClassDesc()
	case tcReference:
		return d.readReference()
	case tcNull:
		return nil, nil
	case tcException:
		return d.readException()
	case tcReset:
		d.Reset()
		return nil, nil
	case tcBlockData:
		return d.readBlockDataShort()
	case tcBlockDataLong:
		return d.readBlockDataLong()
	case tcEndBlockData:
		return nil, nil
	default:
		return nil, fmt.Errorf("unexpected terminal constant: %02x", t)
	}
}

func (d *Decoder) readNewClass() (*Class, error) {
	// newClass:
	//   TC_CLASS classDesc newHandle
	c, err := d.readClassDesc()
	if err != nil {
		return c, fmt.Errorf("d.readClassDesc: %w", err)
	}
	d.o[d.newHandle()] = c
	return c, nil
}

func (d *Decoder) readClassDesc() (*Class, error) {
	// classDesc:
	//   newClassDesc
	//   nullReference
	//   (ClassDesc)prevObject  // an object required to be of type ClassDesc
	t, err := d.r.readByte()
	if err != nil {
		return nil, fmt.Errorf("readByte: %w", err)
	}

	switch t {
	case tcClassDesc:
		cd, err := d.readNewClassDesc()
		return &cd, err
	case tcProxyClassDesc:
		cd, err := d.readNewProxyClassDesc()
		return &cd, err
	case tcNull:
		return nil, nil
	case tcReference:
		o, err := d.readReference()
		if err != nil {
			return nil, fmt.Errorf("d.readReference: %w", err)
		}
		if c, ok := o.(*Class); ok {
			return c, nil
		}
		return nil, fmt.Errorf("invalid reference or non-ClassDesc object")
	default:
		return nil, fmt.Errorf("unexpected terminal constant: %02x", t)
	}
}

func (d *Decoder) readSuperClassDesc() (*Class, error) {
	// superClassDesc:
	//   classDesc
	return d.readClassDesc()
}

func (d *Decoder) readNewClassDesc() (Class, error) {
	// newClassDesc:
	//   TC_CLASSDESC className serialVersionUID newHandle classDescInfo
	//   TC_PROXYCLASSDESC newHandle proxyClassDescInfo
	var err error
	c := Class{}
	c.ClassName, err = d.readString()
	if err != nil {
		return Class{}, fmt.Errorf("d.readNewString: %w", err)
	}
	c.SerialVersionUID, err = d.readSerialVersionUID()
	if err != nil {
		return Class{}, fmt.Errorf("d.readSerialVersionUID: %w", err)
	}
	h := d.newHandle()
	c.Info, err = d.readClassDescInfo()
	if err != nil {
		return Class{}, fmt.Errorf("d.readClassDescInfo: %w", err)
	}
	d.o[h] = &c
	return c, nil
}

func (d *Decoder) readNewProxyClassDesc() (Class, error) {
	// newClassDesc:
	//   TC_CLASSDESC className serialVersionUID newHandle classDescInfo
	//   TC_PROXYCLASSDESC newHandle proxyClassDescInfo
	c := Class{}
	d.o[d.newHandle()] = &c
	return c, ErrNotSupported
}

func (d *Decoder) readClassDescInfo() (ClassDescInfo, error) {
	// classDescInfo:
	//   classDescFlags fields classAnnotation superClassDesc
	var err error
	cdi := ClassDescInfo{}
	cdi.Flags, err = d.readClassDescFlags()
	if err != nil {
		return cdi, fmt.Errorf("d.readClassDescFlags: %w", err)
	}
	cdi.Fields, err = d.readFields()
	if err != nil {
		return cdi, fmt.Errorf("d.readFields: %w", err)
	}
	cdi.ClassAnnotation, err = d.readClassAnnotation()
	if err != nil {
		return cdi, fmt.Errorf("d.readClassAnnotation: %w", err)
	}
	cdi.SuperClassDesc, err = d.readSuperClassDesc()
	if err != nil {
		return cdi, fmt.Errorf("d.readSuperClassDesc: %w", err)
	}
	return cdi, nil
}

func (d *Decoder) readFields() ([]Field, error) {
	// fields:
	//   (short)<count> fieldDesc[count]
	count, err := d.r.readInt16()
	if err != nil {
		return nil, fmt.Errorf("readInt16: %w", err)
	}

	var fields []Field
	for i := 0; i < int(count); i++ {
		f, err := d.readField()
		if err != nil {
			return nil, fmt.Errorf("d.readField: %w", err)
		}
		fields = append(fields, f)
	}
	return fields, nil
}

func (d *Decoder) readField() (Field, error) {
	// fieldDesc:
	//   primitiveDesc
	//   objectDesc
	// primitiveDesc:
	//   prim_typecode fieldName
	// objectDesc:
	//   obj_typecode fieldName className1
	// fieldName:
	//   (utf)
	// className1:
	//   (String)object     // String containing the field's type, in field descriptor format
	f := Field{}
	t, err := d.r.readByte()
	if err != nil {
		return f, fmt.Errorf("readByte: %w", err)
	}
	f.TypeCode = TypeCode(t)
	f.FieldName, err = d.readString()
	if err != nil {
		return f, fmt.Errorf("d.readNewString: %w", err)
	}
	if f.TypeCode.IsObject() {
		c, err := d.readContent()
		if err != nil {
			return f, fmt.Errorf("d.readContent: %w", err)
		}
		var ok bool
		if f.ClassName, ok = c.(string); !ok {
			return f, fmt.Errorf("failed to cast reference to string")
		}
	}
	return f, nil
}

func (d *Decoder) readClassAnnotation() ([]Annotation, error) {
	// classAnnotation:
	//	 endBlockData
	//	 contents endBlockData  // contents written by annotateClass
	return d.readAnnotation()
}

func (d *Decoder) readNewArray() (Array, error) {
	// newArray:
	//   TC_ARRAY classDesc newHandle (int)<size> values[size]
	var err error
	a := Array{}
	a.ClassDesc, err = d.readClassDesc()
	if err != nil {
		return a, fmt.Errorf("d.readClassDesc: %w", err)
	}
	h := d.newHandle()
	count, err := d.r.readInt32()
	if err != nil {
		return a, fmt.Errorf("readInt32: %w", err)
	}
	t, _ := a.ItemType()
	for i := 0; i < int(count); i++ {
		v, err := d.readValue(t)
		if err != nil && !errors.Is(err, ErrNotSupported) {
			return a, fmt.Errorf("d.readValue: %w", err)
		}
		a.Values = append(a.Values, v)
	}
	d.o[h] = a
	return a, nil
}

func (d *Decoder) readNewObject() (Object, error) {
	// newObject:
	//   TC_OBJECT classDesc newHandle classdata[]  // data for each class
	var err error
	o := Object{}
	o.ClassDesc, err = d.readClassDesc()
	if err != nil {
		return o, fmt.Errorf("d.readClassDesc: %w", err)
	}
	h := d.newHandle()
	o.ClassData, err = d.readClassData(o.ClassDesc)
	if err != nil {
		return o, fmt.Errorf("d.readClassData: %w", err)
	}
	d.o[h] = o
	return o, nil
}

func (d *Decoder) readClassData(classDesc *Class) (map[string]ClassData, error) {
	// classdata:
	//   nowrclass                 // SC_SERIALIZABLE & classDescFlag && !(SC_WRITE_METHOD & classDescFlags)
	//   wrclass objectAnnotation  // SC_SERIALIZABLE & classDescFlag && SC_WRITE_METHOD & classDescFlags
	//   externalContents          // SC_EXTERNALIZABLE & classDescFlag && !(SC_BLOCKDATA  & classDescFlags
	//   objectAnnotation          // SC_EXTERNALIZABLE & classDescFlag && SC_BLOCKDATA & classDescFlags
	// nowrclass:
	//   values                    // fields in order of class descriptor
	// wrclass:
	//   nowrclass
	// objectAnnotation:
	//   endBlockData
	//   contents endBlockData   // contents written by writeObject or writeExternal PROTOCOL_VERSION_2.
	// externalContents:         // externalContent written by
	//   externalContent         // writeExternal in PROTOCOL_VERSION_1.
	//   externalContents externalContent
	// externalContent:          // Only parseable by readExternal
	//   (bytes)                 // primitive data
	//   object
	var classDescs []Class
	for cd := classDesc; cd != nil; cd = cd.Info.SuperClassDesc {
		classDescs = append(classDescs, *cd)
	}

	classesData := make(map[string]ClassData)
	for i := len(classDescs) - 1; i >= 0; i-- {
		desc := classDescs[i]
		classData := make(map[string]Value)
		for _, f := range desc.Info.Fields {
			v, err := d.readValue(f.TypeCode)
			if err != nil && !errors.Is(err, ErrNotSupported) {
				return nil, fmt.Errorf("d.readValue: %w", err)
			}
			classData[f.FieldName] = v
		}
		if desc.Info.Flags.IsSerializable() && desc.Info.Flags.HasWriteMethod() {
			contents, err := d.readAnnotation()
			if err != nil {
				return nil, fmt.Errorf("d.readAnnotation: %w", err)
			}
			if len(contents) > 0 {
				classData["[object annotation]"] = contents
			}
		}
		classesData[desc.ClassName] = classData
	}

	return classesData, nil
}

func (d *Decoder) readValue(t TypeCode) (Value, error) {
	switch t {
	case TypeByte:
		return d.r.readInt8()
	case TypeChar:
		b, err := d.r.readInt16()
		return rune(b), err
	case TypeDouble:
		return d.r.readFloat64()
	case TypeFloat:
		return d.r.readFloat32()
	case TypeInteger:
		i, err := d.r.readInt32()
		return int(i), err
	case TypeLong:
		l, err := d.r.readInt64()
		return l, err
	case TypeShort:
		s, err := d.r.readInt16()
		return s, err
	case TypeBoolean:
		b, err := d.r.readByte()
		return b != 0, err
	case TypeArray:
		return d.readContent()
	case TypeObject:
		return d.readContent()
	default:
		return nil, fmt.Errorf("unknown type code: '%c'", t)
	}
}

// blockdata:
//   blockdatashort
//   blockdatalong

// endBlockData:
//   TC_ENDBLOCKDATA

func (d *Decoder) readBlockDataShort() (blockData, error) {
	// blockdatashort:
	//   TC_BLOCKDATA (unsigned byte)<size> (byte)[size]
	l, err := d.r.readByte()
	if err != nil {
		return nil, fmt.Errorf("readByte: %w", err)
	}
	b, err := d.r.readBytes(int64(l))
	if err != nil {
		return nil, fmt.Errorf("readBytes: %w", err)
	}
	return b, nil
}

func (d *Decoder) readBlockDataLong() (blockData, error) {
	// blockdatalong:
	//   TC_BLOCKDATALONG (int)<size> (byte)[size]
	l, err := d.r.readInt32()
	if err != nil {
		return nil, fmt.Errorf("readInt32: %w", err)
	}
	b, err := d.r.readBytes(int64(l))
	if err != nil {
		return nil, fmt.Errorf("readBytes: %w", err)
	}
	return b, nil
}

func (d *Decoder) readNewString() (string, error) {
	// newString:
	//   TC_STRING newHandle (utf)
	//   TC_LONGSTRING newHandle (long-utf)
	h := d.newHandle()
	s, err := d.readString()
	if err != nil {
		return "", fmt.Errorf("d.readString: %w", err)
	}
	d.o[h] = s
	return s, nil
}

func (d *Decoder) readNewLongString() (string, error) {
	// newString:
	//   TC_STRING newHandle (utf)
	//   TC_LONGSTRING newHandle (long-utf)
	h := d.newHandle()
	l, err := d.r.readInt64()
	if err != nil {
		return "", fmt.Errorf("readInt64: %w", err)
	}
	s, err := d.r.readBytes(l)
	if err != nil {
		return "", fmt.Errorf("readBytes: %w", err)
	}
	d.o[h] = string(s)
	return string(s), nil
}

func (d *Decoder) readNewEnum() (Enum, error) {
	// newEnum:
	//   TC_ENUM classDesc newHandle enumConstantName
	// enumConstantName:
	//   (String)object
	var e Enum
	var err error
	e.ClassDesc, err = d.readClassDesc()
	if err != nil {
		return e, fmt.Errorf("d.readClassDesc: %w", err)
	}
	h := d.newHandle()
	c, err := d.readContent()
	if err != nil {
		return e, fmt.Errorf("d.readContent: %w", err)
	}
	var ok bool
	if e.ConstantName, ok = c.(string); !ok {
		return e, fmt.Errorf("failed to cast reference to string")
	}
	d.o[h] = e
	return e, nil
}

func (d *Decoder) newHandle() handle {
	h := d.h
	d.h++
	return h
}

func (d *Decoder) readAnnotation() ([]Annotation, error) {
	var annotations []Annotation
	for {
		content, err := d.readContent()
		if err != nil {
			return nil, err
		}
		if content == nil {
			break
		}
		annotations = append(annotations, content)
	}
	return annotations, nil
}

func (d *Decoder) readReference() (Content, error) {
	// prevObject:
	//   TC_REFERENCE (int)handle
	h, err := d.r.readInt32()
	if err != nil {
		return nil, fmt.Errorf("readInt32: %w", err)
	}
	return d.o[handle(h)], err
}

func (d *Decoder) readException() (Object, error) {
	// exception:
	//   TC_EXCEPTION reset (Throwable)object reset
	d.Reset()
	e, err := d.readNewObject()
	d.Reset()
	return e, err
}

func (d *Decoder) readSerialVersionUID() (SerialVersionUID, error) {
	u, err := d.r.readInt64()
	if err != nil {
		return 0, fmt.Errorf("readInt64: %w", err)
	}
	return SerialVersionUID(u), nil
}

func (d *Decoder) readClassDescFlags() (ClassDescFlags, error) {
	f, err := d.r.readByte()
	if err != nil {
		return 0, fmt.Errorf("readByte: %w", err)
	}
	return ClassDescFlags(f), nil
}

func (d *Decoder) readString() (string, error) {
	l, err := d.r.readUint16()
	if err != nil {
		return "", fmt.Errorf("readUint16: %w", err)
	}
	s, err := d.r.readBytes(int64(l))
	if err != nil {
		return "", fmt.Errorf("readBytes: %w", err)
	}
	return string(s), nil
}

type blockData []byte
type handle int

const baseHandleValue = 0x7e0000

const (
	tcNull           = 0x70
	tcReference      = 0x71
	tcClassDesc      = 0x72
	tcObject         = 0x73
	tcString         = 0x74
	tcArray          = 0x75
	tcClass          = 0x76
	tcBlockData      = 0x77
	tcEndBlockData   = 0x78
	tcReset          = 0x79
	tcBlockDataLong  = 0x7A
	tcException      = 0x7B
	tcLongString     = 0x7C
	tcProxyClassDesc = 0x7D
	tcEnum           = 0x7E
)
