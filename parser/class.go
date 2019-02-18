package parser

import (
	"fmt"
	"io"
)

type Class struct {
	classFile *ClassFile
}

func ReadClass(reader io.Reader) (*Class, error) {
	classFile, err := Read(reader)
	if err != nil {
		return nil, err
	}
	return &Class{classFile}, nil
}

func (c *Class) Name() string {
	return c.classFile.ConstantPool.GetClass(c.classFile.ThisClass)
}

func (c *Class) SuperClassName() string {
	return c.classFile.ConstantPool.GetClass(c.classFile.SuperClass)
}

func (c *Class) Interfaces() []string {
	interfaces := make([]string, len(c.classFile.Interfaces))
	for _, intf := range c.classFile.Interfaces {
		interfaces = append(interfaces, c.classFile.ConstantPool.GetClass(intf))
	}
	return interfaces
}

func (c *Class) AccessFlags() AccessFlags {
	return c.classFile.AccessFlags
}

func (c *Class) IsInterface() bool {
	return c.classFile.AccessFlags.Interface()
}

func (c *Class) IsEnum() bool {
	return c.classFile.AccessFlags.Enum()
}

func (c *Class) IsAnnotation() bool {
	return c.classFile.AccessFlags.Annotation()
}

func (c *Class) IsAbstract() bool {
	return c.classFile.AccessFlags.Abstract()
}

func (c Class) String() string {
	return fmt.Sprintf("Class[name=%s, super=%s, interfaces=%v, access=%v]", c.Name(), c.SuperClassName(), c.Interfaces(), c.AccessFlags())
}
