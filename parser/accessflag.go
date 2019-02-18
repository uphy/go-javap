package parser

import (
	"fmt"
	"strings"
)

type AccessFlags uint16

const (
	AccessPublic     = 0x0001
	AccessFinal      = 0x0010
	AccessSuper      = 0x0020
	AccessInterface  = 0x0200
	AccessAbstract   = 0x0400
	AccessSynthetic  = 0x1000
	AccessAnnotation = 0x2000
	AccessEnum       = 0x4000
)

func (a AccessFlags) Public() bool {
	return a.is(AccessPublic)
}

func (a AccessFlags) Final() bool {
	return a.is(AccessFinal)
}

func (a AccessFlags) Super() bool {
	return a.is(AccessSuper)
}

func (a AccessFlags) Interface() bool {
	return a.is(AccessInterface)
}

func (a AccessFlags) Abstract() bool {
	return a.is(AccessAbstract)
}

func (a AccessFlags) Synthetic() bool {
	return a.is(AccessSynthetic)
}

func (a AccessFlags) Annotation() bool {
	return a.is(AccessAnnotation)
}

func (a AccessFlags) Enum() bool {
	return a.is(AccessAnnotation)
}

func (a AccessFlags) String() string {
	flags := make([]string, 0)
	if a.Public() {
		flags = append(flags, "public")
	}
	if a.Final() {
		flags = append(flags, "final")
	}
	if a.Super() {
		flags = append(flags, "super")
	}
	if a.Interface() {
		flags = append(flags, "interface")
	}
	if a.Abstract() {
		flags = append(flags, "abstract")
	}
	if a.Synthetic() {
		flags = append(flags, "synthetic")
	}
	if a.Annotation() {
		flags = append(flags, "annotation")
	}
	if a.Enum() {
		flags = append(flags, "enum")
	}
	return fmt.Sprintf("AccessFlags[%s]", strings.Join(flags, ", "))
}

func (a AccessFlags) is(n uint16) bool {
	return (uint16(a) & n) != 0
}
