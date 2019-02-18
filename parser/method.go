package parser

import (
	"fmt"
	"strings"
)

type (
	MethodInfo struct {
		AccessFlags     MethodAccessFlags
		NameIndex       uint16
		DescriptorIndex uint16
		Attributes      []AttributeInfo
	}

	MethodAccessFlags uint16
)

const (
	MethodAccessPublic       = 0x0001
	MethodAccessPrivate      = 0x0002
	MethodAccessProtected    = 0x0004
	MethodAccessStatic       = 0x0008
	MethodAccessFinal        = 0x0010
	MethodAccessSynchronized = 0x0020
	MethodAccessBridge       = 0x0040
	MethodAccessVarArgs      = 0x0080
	MethodAccessNative       = 0x0100
	MethodAccessAbstract     = 0x0400
	MethodAccessStrict       = 0x0800
	MethodAccessSynthetic    = 0x1000
)

func (f MethodInfo) String() string {
	return fmt.Sprintf("Method[flags=%s, nameIndex=%d, descriptorIndex=%d]", f.AccessFlags, f.NameIndex, f.DescriptorIndex)
}

func (a MethodAccessFlags) Public() bool {
	return a.is(MethodAccessPublic)
}

func (a MethodAccessFlags) Private() bool {
	return a.is(MethodAccessPrivate)
}

func (a MethodAccessFlags) Protected() bool {
	return a.is(MethodAccessProtected)
}

func (a MethodAccessFlags) Static() bool {
	return a.is(MethodAccessStatic)
}

func (a MethodAccessFlags) Final() bool {
	return a.is(MethodAccessFinal)
}

func (a MethodAccessFlags) Synchronized() bool {
	return a.is(MethodAccessSynchronized)
}

func (a MethodAccessFlags) Bridge() bool {
	return a.is(MethodAccessBridge)
}

func (a MethodAccessFlags) VarArgs() bool {
	return a.is(MethodAccessVarArgs)
}

func (a MethodAccessFlags) Native() bool {
	return a.is(MethodAccessNative)
}

func (a MethodAccessFlags) Abstract() bool {
	return a.is(MethodAccessAbstract)
}

func (a MethodAccessFlags) Strict() bool {
	return a.is(MethodAccessStrict)
}

func (a MethodAccessFlags) Synthetic() bool {
	return a.is(MethodAccessSynthetic)
}

func (a MethodAccessFlags) String() string {
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
	if a.Synchronized() {
		flags = append(flags, "synchronized")
	}
	if a.Bridge() {
		flags = append(flags, "bridge")
	}
	if a.VarArgs() {
		flags = append(flags, "varargs")
	}
	if a.Native() {
		flags = append(flags, "native")
	}
	if a.Abstract() {
		flags = append(flags, "abstract")
	}
	if a.Strict() {
		flags = append(flags, "strict")
	}
	if a.Synthetic() {
		flags = append(flags, "synthetic")
	}
	return fmt.Sprintf("AccessFlags[%s]", strings.Join(flags, ", "))
}

func (a MethodAccessFlags) is(n uint16) bool {
	return (uint16(a) & n) != 0
}
