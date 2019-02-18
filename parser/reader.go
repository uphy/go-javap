package parser

import (
	"bufio"
	"io"
)

type Reader struct {
	reader *bufio.Reader
}

func NewReader(r io.Reader) *Reader {
	return &Reader{bufio.NewReader(r)}
}

func (r *Reader) Read8() (uint8, error) {
	return r.reader.ReadByte()
}

func (r *Reader) Read16() (uint16, error) {
	n1, err := r.Read8()
	if err != nil {
		return 0, err
	}
	n2, err := r.Read8()
	if err != nil {
		return 0, err
	}
	return uint16(n1)<<8 | uint16(n2), nil
}

func (r *Reader) Read32() (uint32, error) {
	n1, err := r.Read16()
	if err != nil {
		return 0, err
	}
	n2, err := r.Read16()
	if err != nil {
		return 0, err
	}
	return uint32(n1)<<16 | uint32(n2), nil
}

func (r *Reader) Read64() (uint64, error) {
	n1, err := r.Read32()
	if err != nil {
		return 0, err
	}
	n2, err := r.Read32()
	if err != nil {
		return 0, err
	}
	return uint64(n1)<<32 | uint64(n2), nil
}

func (r *Reader) ReadBytes(bytes []byte) (int, error) {
	read := 0
	for {
		n, err := r.reader.Read(bytes[read:])
		if err != nil {
			return 0, err
		}

		read += n
		if read >= len(bytes) {
			return read, nil
		}
	}
}

func (r *Reader) ReadModifiedUTF8(length uint16) (string, error) {
	var index uint16
	runes := make([]rune, length)
	for {
		n1, err := r.Read8()
		if err != nil {
			return "", err
		}
		var c rune
		if n1 <= 0x7F {
			c = rune(n1)
		} else {
			n2, err := r.Read8()
			if err != nil {
				return "", err
			}
			w := (uint16(n1) << 8) | uint16(n2)
			if w <= 0x7FF {
				c = rune(((uint16(n1) & 0x1F) << 6) | (uint16(n2) & 0x3F))
			} else {
				n3, err := r.Read8()
				if err != nil {
					return "", err
				}
				c = rune(((uint32(n1) & 0x0F) << 12) | ((uint32(n2) & 0x3F) << 6) | (uint32(n3) & 0x3F))
			}
		}
		runes[index] = c
		index++
		if index == length {
			break
		}
	}
	return string(runes), nil
}
