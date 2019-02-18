package parser

import (
	"bufio"
	"bytes"
	"testing"
)

func TestReader_Read8(t *testing.T) {
	type fields struct {
		reader *bufio.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		want    uint8
		wantErr bool
	}{
		{
			fields: fields{
				reader: bufio.NewReader(bytes.NewReader([]byte{0})),
			},
			want: 0,
		},
		{
			fields: fields{
				reader: bufio.NewReader(bytes.NewReader([]byte{0x10})),
			},
			want: 0x10,
		},
		{
			fields: fields{
				reader: bufio.NewReader(bytes.NewReader([]byte{})),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Reader{
				reader: tt.fields.reader,
			}
			got, err := r.Read8()
			if (err != nil) != tt.wantErr {
				t.Errorf("Reader.Read8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Reader.Read8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReader_Read16(t *testing.T) {
	type fields struct {
		reader *bufio.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		want    uint16
		wantErr bool
	}{
		{
			fields: fields{
				reader: bufio.NewReader(bytes.NewReader([]byte{0x01, 0x02})),
			},
			want: 0x0102,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Reader{
				reader: tt.fields.reader,
			}
			got, err := r.Read16()
			if (err != nil) != tt.wantErr {
				t.Errorf("Reader.Read16() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Reader.Read16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReader_Read32(t *testing.T) {
	type fields struct {
		reader *bufio.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		want    uint32
		wantErr bool
	}{
		{
			fields: fields{
				reader: bufio.NewReader(bytes.NewReader([]byte{0x01, 0x02, 0x03, 0x04})),
			},
			want: 0x01020304,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Reader{
				reader: tt.fields.reader,
			}
			got, err := r.Read32()
			if (err != nil) != tt.wantErr {
				t.Errorf("Reader.Read32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Reader.Read32() = %v, want %v", got, tt.want)
			}
		})
	}
}
