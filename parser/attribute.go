package parser

import (
	"encoding/hex"
	"fmt"
)

type AttributeInfo struct {
	NameIndex uint16
	Attribute []byte
}

func (a AttributeInfo) String() string {
	return fmt.Sprintf("AttributeInfo[nameIndex=%d, attribute=%s]", a.NameIndex, hex.EncodeToString(a.Attribute))
}

func readAttribute(r *Reader) ([]AttributeInfo, error) {
	attributesCount, _ := r.Read16()
	attributes := make([]AttributeInfo, attributesCount)
	for j := uint16(0); j < attributesCount; j++ {
		attributeNameIndex, _ := r.Read16()
		attributeLength, _ := r.Read32()
		info := make([]byte, attributeLength)
		r.ReadBytes(info)
		attributes = append(attributes, AttributeInfo{attributeNameIndex, info})
	}
	return nil, nil
}
