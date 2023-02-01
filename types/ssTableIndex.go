package types

import (
	"bytes"
	"encoding/binary"
)

type Index struct {
	KeySize uint64
	Key     string
	Offset  uint64
}

func (index Index) Serialize() []byte {
	var serializedIndex = new(bytes.Buffer)

	binary.Write(serializedIndex, binary.LittleEndian, []byte(index.Key))
	binary.Write(serializedIndex, binary.LittleEndian, index.Offset)
	return serializedIndex.Bytes()
}

func DeserializeIndex(serializedIndex []byte) Index {
	return Index{}
}
