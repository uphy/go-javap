package parser

import (
	"fmt"
	"strings"
)

type (
	FieldInfo struct {
		AccessFlags     FieldAccessFlags
		NameIndex       uint16
		DescriptorIndex uint16
		Attributes      []AttributeInfo
	}

	FieldAccessFlags uint16
)

const (
	FieldAccessPublic    = 0x0001
	FieldAccessPrivate   = 0x0002
	FieldAccessProtected = 0x0004
	FieldAccessStatic    = 0x0008
	FieldAccessFinal     = 0x0010
	FieldAccessVolatile  = 0x0040
	FieldAccessTransient = 0x0080
	FieldAccessSynthetic = 0x1000
	FieldAccessEnum      = 0x4000
)

func (f FieldInfo) String() string {
	return fmt.Sprintf("Field[flags=%s, nameIndex=%d, descriptorIndex=%d]", f.AccessFlags, f.NameIndex, f.DescriptorIndex)
}

func (a FieldAccessFlags) Public() bool {
	return a.is(FieldAccessPublic)
}

func (a FieldAccessFlags) Private() bool {
	return a.is(FieldAccessPrivate)
}

func (a FieldAccessFlags) Protected() bool {
	return a.is(FieldAccessProtected)
}

func (a FieldAccessFlags) Static() bool {
	return a.is(FieldAccessStatic)
}

func (a FieldAccessFlags) Final() bool {
	return a.is(FieldAccessFinal)
}

func (a FieldAccessFlags) Volatile() bool {
	return a.is(FieldAccessVolatile)
}

func (a FieldAccessFlags) Transient() bool {
	return a.is(FieldAccessTransient)
}

func (a FieldAccessFlags) Synthetic() bool {
	return a.is(FieldAccessSynthetic)
}

func (a FieldAccessFlags) Enum() bool {
	return a.is(FieldAccessEnum)
}

func (a FieldAccessFlags) String() string {
	flags := make([]string, 0)
	if a.Public() {
		flags = append(flags, "public")
	}
	if a.Private() {
		flags = append(flags, "private")
	}
	if a.Protected() {
		flags = append(flags, "protected")
	}
	if a.Static() {
		flags = append(flags, "static")
	}
	if a.Final() {
		flags = append(flags, "final")
	}
	if a.Volatile() {
		flags = append(flags, "volatile")
	}
	if a.Transient() {
		flags = append(flags, "transient")
	}
	if a.Synthetic() {
		flags = append(flags, "synthetic")
	}
	if a.Enum() {
		flags = append(flags, "enum")
	}
	return fmt.Sprintf("AccessFlags[%s]", strings.Join(flags, ", "))
}

func (a FieldAccessFlags) is(n uint16) bool {
	return (uint16(a) & n) != 0
}
