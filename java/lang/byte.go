package lang

type Byte struct {
	Value int8 `java:"value"`
}

func (b *Byte) ByteValue() byte {
	return byte(b.Value)
}
