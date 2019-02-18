package parser

import "fmt"

const (
	ConstantUtf8               = 0x01
	ConstantInteger            = 0x03
	ConstantFloat              = 0x04
	ConstantLong               = 0x05
	ConstantDouble             = 0x06
	ConstantClass              = 0x07
	ConstantString             = 0x08
	ConstantFieldref           = 0x09
	ConstantMethodref          = 0x0A
	ConstantInterfaceMethodref = 0x0B
	ConstantNameAndType        = 0x0C
	ConstantMethodHandle       = 0x0F
	ConstantMethodType         = 0x10
	ConstantInvokeDynamic      = 0x12
	ConstantModule             = 0x13
	ConstantPackage            = 0x14

	MethodHandleRefGetField         MethodHandleRef = 0x01
	MethodHandleRefGetStatic        MethodHandleRef = 0x02
	MethodHandleRefPutField         MethodHandleRef = 0x03
	MethodHandleRefPutStatic        MethodHandleRef = 0x04
	MethodHandleRefInvokeVirtual    MethodHandleRef = 0x05
	MethodHandleRefInvokeStatic     MethodHandleRef = 0x06
	MethodHandleRefInvokeSpecial    MethodHandleRef = 0x07
	MethodHandleRefNewInvokeSpecial MethodHandleRef = 0x08
	MethodHandleRefInvokeInterface  MethodHandleRef = 0x09
)

type (
	ConstantPool     []ConstantInfo
	ConstantUtf8Info struct {
		Bytes []byte
	}
	ConstantIntegerInfo struct {
		Value int32
	}
	ConstantFloatInfo struct {
		Value float32
	}
	ConstantLongInfo struct {
		Value int64
	}
	ConstantDoubleInfo struct {
		Value float64
	}
	ConstantInfo interface {
		String() string
	}
	ConstantClassInfo struct {
		NameIndex uint16
	}
	ConstantMethodrefInfo struct {
		ClassIndex       uint16
		NameAndTypeIndex uint16
	}
	ConstantInterfaceMethodrefInfo struct {
		ClassIndex       uint16
		NameAndTypeIndex uint16
	}
	ConstantNameAndTypeInfo struct {
		NameIndex       uint16
		DescriptorIndex uint16
	}
	ConstantFieldrefInfo struct {
		ClassIndex       uint16
		NameAndTypeIndex uint16
	}
	ConstantStringInfo struct {
		StringIndex uint16
	}
	MethodHandleRef          uint16
	ConstantMethodHandleInfo struct {
		ReferenceKind  MethodHandleRef
		ReferenceIndex uint16
	}
	ConstantMethodTypeInfo struct {
		DescriptorIndex uint16
	}
	ConstantInvokeDynamicInfo struct {
		BootstrapMethodAttrIndex uint16
		NameAndTypeIndex         uint16
	}
	ConstantModuleInfo struct {
		NameIndex uint16
	}
	ConstantPackageInfo struct {
		NameIndex uint16
	}
)

func (p ConstantPool) get(index uint16) ConstantInfo {
	return p[index-1]
}

func (p ConstantPool) GetUTF8(index uint16) string {
	if index == 0 {
		return ""
	}
	info := p.get(index)
	if i, ok := info.(ConstantUtf8Info); ok {
		return string(i.Bytes)
	}
	return ""
}

func (p ConstantPool) GetString(index uint16) string {
	if index == 0 {
		return ""
	}
	info := p.get(index)
	if i, ok := info.(ConstantStringInfo); ok {
		return p.GetUTF8(i.StringIndex)
	}
	return ""
}

func (p ConstantPool) GetClass(index uint16) string {
	if index == 0 {
		return ""
	}
	info := p.get(index)
	if i, ok := info.(ConstantClassInfo); ok {
		return p.GetUTF8(i.NameIndex)
	}
	return ""
}

func (c ConstantClassInfo) String() string {
	return fmt.Sprintf("Class[nameIndex=%d]", c.NameIndex)
}

func (c ConstantIntegerInfo) String() string {
	return fmt.Sprintf("Integer[value=%d]", c.Value)
}

func (c ConstantFloatInfo) String() string {
	return fmt.Sprintf("Float[value=%f]", c.Value)
}

func (c ConstantLongInfo) String() string {
	return fmt.Sprintf("Long[value=%d]", c.Value)
}

func (c ConstantDoubleInfo) String() string {
	return fmt.Sprintf("Double[value=%f]", c.Value)
}

func (c ConstantMethodrefInfo) String() string {
	return fmt.Sprintf("Methodref[classIndex=%d, nameAndTypeIndex=%d]", c.ClassIndex, c.NameAndTypeIndex)
}

func (c ConstantInterfaceMethodrefInfo) String() string {
	return fmt.Sprintf("InterfaceMethodrefInfo[classIndex=%d, nameAndTypeIndex=%d]", c.ClassIndex, c.NameAndTypeIndex)
}

func (c ConstantUtf8Info) String() string {
	return fmt.Sprintf("UTF[%s]", string(c.Bytes))
}

func (c ConstantNameAndTypeInfo) String() string {
	return fmt.Sprintf("NameAndType[nameIndex=%d, descriptorIndex=%d]", c.NameIndex, c.DescriptorIndex)
}

func (c ConstantFieldrefInfo) String() string {
	return fmt.Sprintf("Fieldref[classIndex=%d, nameAndTypeIndex=%d]", c.ClassIndex, c.NameAndTypeIndex)
}

func (c ConstantStringInfo) String() string {
	return fmt.Sprintf("String[stringIndex=%d]", c.StringIndex)
}

func (c ConstantMethodHandleInfo) String() string {
	return fmt.Sprintf("MethodHandle[referenceKind=%s, referenceIndex=%d]", c.ReferenceKind, c.ReferenceIndex)
}

func (c MethodHandleRef) String() string {
	switch c {
	case MethodHandleRefGetField:
		return "getField"
	case MethodHandleRefGetStatic:
		return "getStatic"
	case MethodHandleRefPutField:
		return "putField"
	case MethodHandleRefPutStatic:
		return "putStatic"
	case MethodHandleRefInvokeVirtual:
		return "invokeVirtual"
	case MethodHandleRefInvokeStatic:
		return "invokeStatic"
	case MethodHandleRefInvokeSpecial:
		return "invokeSpecial"
	case MethodHandleRefNewInvokeSpecial:
		return "newInvokeSpecial"
	case MethodHandleRefInvokeInterface:
		return "invokeInterface"
	}
	return "unknown"
}

func (c ConstantMethodTypeInfo) String() string {
	return fmt.Sprintf("MethodType[descriptorIndex=%d]", c.DescriptorIndex)
}

func (c ConstantInvokeDynamicInfo) String() string {
	return fmt.Sprintf("InvokeDynamic[bootstrapMethodAttrIndex=%d, nameAndTypeIndex=%d]", c.BootstrapMethodAttrIndex, c.NameAndTypeIndex)
}

func (c ConstantModuleInfo) String() string {
	return fmt.Sprintf("Module[nameIndex=%d]", c.NameIndex)
}

func (c ConstantPackageInfo) String() string {
	return fmt.Sprintf("Package[nameIndex=%d]", c.NameIndex)
}
