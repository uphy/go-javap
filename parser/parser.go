package parser

import (
	"fmt"
	"io"
	"math"
)

type (
	ClassFile struct {
		MinorVersion uint16
		MajorVersion uint16
		ConstantPool ConstantPool
		AccessFlags  AccessFlags
		ThisClass    uint16
		SuperClass   uint16
		Interfaces   []uint16
		Fields       []FieldInfo
		Methods      []MethodInfo
		Attributes   []AttributeInfo
	}
)

func Read(reader io.Reader) (*ClassFile, error) {
	r := NewReader(reader)
	magic, err := r.Read32()
	if err != nil {
		return nil, err
	}
	if magic != 0xCAFEBABE {
		return nil, fmt.Errorf("unexpected magic: %08X", magic)
	}
	c := new(ClassFile)
	c.MinorVersion, _ = r.Read16()
	c.MajorVersion, _ = r.Read16()
	constantPoolCount, _ := r.Read16()
	for i := uint16(0); i < constantPoolCount-1; i++ {
		constantType, _ := r.Read8()
		var info ConstantInfo
		double := false
		switch constantType {
		case ConstantMethodref:
			classIndex, _ := r.Read16()
			nameAndTypeIndex, _ := r.Read16()
			info = ConstantMethodrefInfo{classIndex, nameAndTypeIndex}
		case ConstantInterfaceMethodref:
			classIndex, _ := r.Read16()
			nameAndTypeIndex, _ := r.Read16()
			info = ConstantInterfaceMethodrefInfo{classIndex, nameAndTypeIndex}
		case ConstantClass:
			nameIndex, _ := r.Read16()
			info = ConstantClassInfo{nameIndex}
		case ConstantUtf8:
			length, _ := r.Read16()
			bytes := make([]byte, length)
			r.ReadBytes(bytes)
			info = ConstantUtf8Info{bytes}
		case ConstantNameAndType:
			nameIndex, _ := r.Read16()
			descriptorIndex, _ := r.Read16()
			info = ConstantNameAndTypeInfo{nameIndex, descriptorIndex}
		case ConstantFieldref:
			classIndex, _ := r.Read16()
			nameAndTypeIndex, _ := r.Read16()
			info = ConstantFieldrefInfo{classIndex, nameAndTypeIndex}
		case ConstantString:
			stringIndex, _ := r.Read16()
			info = ConstantStringInfo{stringIndex}
		case ConstantInteger:
			value, _ := r.Read32()
			info = ConstantIntegerInfo{int32(value)}
		case ConstantLong:
			value, _ := r.Read64()
			info = ConstantLongInfo{int64(value)}
			double = true
		case ConstantFloat:
			bitsuint, _ := r.Read32()
			bits := int32(bitsuint)
			var s int32
			if (bits >> 31) == 0 {
				s = 1
			} else {
				s = -1
			}
			e := (bits >> 23) & 0x7ff
			var m int32
			if e == 0 {
				m = (bits & 0x7fffff) << 1
			} else {
				m = (bits & 0x7fffff) | 0x800000
			}
			d := float64(s) * float64(m) * math.Pow(math.E, -1075.)
			info = ConstantFloatInfo{float32(d)}
		case ConstantDouble:
			double = true
			bitsuint, _ := r.Read64()
			bits := int64(bitsuint)
			var s int64
			if (bits >> 63) == 0 {
				s = 1
			} else {
				s = -1
			}
			e := (bits >> 52) & 0x7ff
			var m int64
			if e == 0 {
				m = (bits & 0xfffffffffffff) << 1
			} else {
				m = (bits & 0xfffffffffffff) | 0x10000000000000
			}
			d := float64(s) * float64(m) * math.Pow(math.E, -1075.)
			info = ConstantDoubleInfo{d}
		case ConstantMethodHandle:
			kind, _ := r.Read8()
			index, _ := r.Read16()
			info = ConstantMethodHandleInfo{MethodHandleRef(kind), index}
		case ConstantMethodType:
			descriptorIndex, _ := r.Read16()
			info = ConstantMethodTypeInfo{descriptorIndex}
		case ConstantInvokeDynamic:
			bootstrapMethodAttrIndex, _ := r.Read16()
			nameAndTypeIndex, _ := r.Read16()
			info = ConstantInvokeDynamicInfo{bootstrapMethodAttrIndex, nameAndTypeIndex}
		case ConstantModule:
			nameIndex, _ := r.Read16()
			info = ConstantModuleInfo{nameIndex}
		case ConstantPackage:
			nameIndex, _ := r.Read16()
			info = ConstantPackageInfo{nameIndex}
		default:
			return nil, fmt.Errorf("unsupported constant pool type: 0x%02X(index=%d, constantPoolCount=%d)", constantType, i, constantPoolCount)
		}
		c.ConstantPool = append(c.ConstantPool, info)
		if double {
			c.ConstantPool = append(c.ConstantPool, info)
			i++
		}
	}
	accessFlags, _ := r.Read16()
	c.AccessFlags = AccessFlags(accessFlags)
	c.ThisClass, _ = r.Read16()
	c.SuperClass, _ = r.Read16()
	interfacesCount, _ := r.Read16()
	for i := uint16(0); i < interfacesCount; i++ {
		intf, _ := r.Read16()
		c.Interfaces = append(c.Interfaces, intf)
	}

	fieldsCount, _ := r.Read16()
	for i := uint16(0); i < fieldsCount; i++ {
		flags, _ := r.Read16()
		nameIndex, _ := r.Read16()
		descriptorIndex, _ := r.Read16()
		attributes, _ := readAttribute(r)
		c.Fields = append(c.Fields, FieldInfo{FieldAccessFlags(flags), nameIndex, descriptorIndex, attributes})
	}

	methodsCount, _ := r.Read16()
	for i := uint16(0); i < methodsCount; i++ {
		flags, _ := r.Read16()
		nameIndex, _ := r.Read16()
		descriptorIndex, _ := r.Read16()
		attributes, _ := readAttribute(r)
		c.Methods = append(c.Methods, MethodInfo{MethodAccessFlags(flags), nameIndex, descriptorIndex, attributes})
	}
	c.Attributes, _ = readAttribute(r)
	return c, nil
}
