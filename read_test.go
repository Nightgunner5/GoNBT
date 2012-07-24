package nbt

import (
	"os"
	"reflect"
	"testing"
)

// The main type to be decoded must have the same name as the
// first non-blank compound tag name which in the case of bigtest.nbt is Level
type Level struct {
	// Nested type structures are analogous to compound tags
	Nested     Nested `nbt:"nested compound test"`
	ByteTest   byte
	IntTest    int32
	StringTest string
	// Struct field tags behave much in the same way as the json
	// or xml encoding packages in golang, the name of the field
	// may be specified if the field name cannot be represented legally in go
	ListTestLong     []int64 `nbt:"listTest (long)"`
	DoubleTest       float64
	FloatTest        float32
	LongTest         int64
	ListTestCompound []ListTest `nbt:"listtest (compound)"`
	ByteArrayTest    []byte     `nbt:"bytearraytest (the first 1000 values of (n*n*255+n*7)%100, starting with n=0 (0, 62, 34, 16, 8, ...))"`
	ShortTest        int16
}

type Nested struct {
	Egg Egg
	Ham Ham
}

type Egg struct {
	Name  string
	Value float32
}

type Ham struct {
	Name  string
	Value float32
}

type ListTest struct {
	CreatedOn int64 `nbt:"created-on"`
	Name      string
}

func TestBigTest(t *testing.T) {
	f, err := os.Open("bigtest.nbt")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()

	var l Level
	Read(f, &l)

	correct := Level{
		Nested: Nested{
			Egg: Egg{Name: "Eggbert", Value: 0.5},
			Ham: Ham{Name: "Hampus", Value: 0.75},
		},
		ByteTest:     0x7f,
		IntTest:      2147483647,
		StringTest:   "HELLO WORLD THIS IS A TEST STRING ÅÄÖ!",
		ListTestLong: []int64{11, 12, 13, 14, 15},
		DoubleTest:   0.4931287132182315,
		FloatTest:    0.49823147,
		LongTest:     9223372036854775807,
		ListTestCompound: []ListTest{
			{CreatedOn: 1264099775885, Name: "Compound tag #0"},
			{CreatedOn: 1264099775885, Name: "Compound tag #1"},
		},
		ByteArrayTest: []byte{
			0x0, 0x3e, 0x22, 0x10, 0x8, 0xa, 0x16, 0x2c, 0x4c, 0x12, 0x46, 0x20, 0x4, 0x56, 0x4e,
			0x50, 0x5c, 0xe, 0x2e, 0x58, 0x28, 0x2, 0x4a, 0x38, 0x30, 0x32, 0x3e, 0x54, 0x10, 0x3a,
			0xa, 0x48, 0x2c, 0x1a, 0x12, 0x14, 0x20, 0x36, 0x56, 0x1c, 0x50, 0x2a, 0xe, 0x60, 0x58,
			0x5a, 0x2, 0x18, 0x38, 0x62, 0x32, 0xc, 0x54, 0x42, 0x3a, 0x3c, 0x48, 0x5e, 0x1a, 0x44,
			0x14, 0x52, 0x36, 0x24, 0x1c, 0x1e, 0x2a, 0x40, 0x60, 0x26, 0x5a, 0x34, 0x18, 0x6, 0x62,
			0x0, 0xc, 0x22, 0x42, 0x8, 0x3c, 0x16, 0x5e, 0x4c, 0x44, 0x46, 0x52, 0x4, 0x24, 0x4e,
			0x1e, 0x5c, 0x40, 0x2e, 0x26, 0x28, 0x34, 0x4a, 0x6, 0x30, 0x0, 0x3e, 0x22, 0x10, 0x8,
			0xa, 0x16, 0x2c, 0x4c, 0x12, 0x46, 0x20, 0x4, 0x56, 0x4e, 0x50, 0x5c, 0xe, 0x2e, 0x58,
			0x28, 0x2, 0x4a, 0x38, 0x30, 0x32, 0x3e, 0x54, 0x10, 0x3a, 0xa, 0x48, 0x2c, 0x1a, 0x12,
			0x14, 0x20, 0x36, 0x56, 0x1c, 0x50, 0x2a, 0xe, 0x60, 0x58, 0x5a, 0x2, 0x18, 0x38, 0x62,
			0x32, 0xc, 0x54, 0x42, 0x3a, 0x3c, 0x48, 0x5e, 0x1a, 0x44, 0x14, 0x52, 0x36, 0x24, 0x1c,
			0x1e, 0x2a, 0x40, 0x60, 0x26, 0x5a, 0x34, 0x18, 0x6, 0x62, 0x0, 0xc, 0x22, 0x42, 0x8,
			0x3c, 0x16, 0x5e, 0x4c, 0x44, 0x46, 0x52, 0x4, 0x24, 0x4e, 0x1e, 0x5c, 0x40, 0x2e, 0x26,
			0x28, 0x34, 0x4a, 0x6, 0x30, 0x0, 0x3e, 0x22, 0x10, 0x8, 0xa, 0x16, 0x2c, 0x4c, 0x12,
			0x46, 0x20, 0x4, 0x56, 0x4e, 0x50, 0x5c, 0xe, 0x2e, 0x58, 0x28, 0x2, 0x4a, 0x38, 0x30,
			0x32, 0x3e, 0x54, 0x10, 0x3a, 0xa, 0x48, 0x2c, 0x1a, 0x12, 0x14, 0x20, 0x36, 0x56, 0x1c,
			0x50, 0x2a, 0xe, 0x60, 0x58, 0x5a, 0x2, 0x18, 0x38, 0x62, 0x32, 0xc, 0x54, 0x42, 0x3a,
			0x3c, 0x48, 0x5e, 0x1a, 0x44, 0x14, 0x52, 0x36, 0x24, 0x1c, 0x1e, 0x2a, 0x40, 0x60, 0x26,
			0x5a, 0x34, 0x18, 0x6, 0x62, 0x0, 0xc, 0x22, 0x42, 0x8, 0x3c, 0x16, 0x5e, 0x4c, 0x44,
			0x46, 0x52, 0x4, 0x24, 0x4e, 0x1e, 0x5c, 0x40, 0x2e, 0x26, 0x28, 0x34, 0x4a, 0x6, 0x30,
			0x0, 0x3e, 0x22, 0x10, 0x8, 0xa, 0x16, 0x2c, 0x4c, 0x12, 0x46, 0x20, 0x4, 0x56, 0x4e,
			0x50, 0x5c, 0xe, 0x2e, 0x58, 0x28, 0x2, 0x4a, 0x38, 0x30, 0x32, 0x3e, 0x54, 0x10, 0x3a,
			0xa, 0x48, 0x2c, 0x1a, 0x12, 0x14, 0x20, 0x36, 0x56, 0x1c, 0x50, 0x2a, 0xe, 0x60, 0x58,
			0x5a, 0x2, 0x18, 0x38, 0x62, 0x32, 0xc, 0x54, 0x42, 0x3a, 0x3c, 0x48, 0x5e, 0x1a, 0x44,
			0x14, 0x52, 0x36, 0x24, 0x1c, 0x1e, 0x2a, 0x40, 0x60, 0x26, 0x5a, 0x34, 0x18, 0x6, 0x62,
			0x0, 0xc, 0x22, 0x42, 0x8, 0x3c, 0x16, 0x5e, 0x4c, 0x44, 0x46, 0x52, 0x4, 0x24, 0x4e,
			0x1e, 0x5c, 0x40, 0x2e, 0x26, 0x28, 0x34, 0x4a, 0x6, 0x30, 0x0, 0x3e, 0x22, 0x10, 0x8,
			0xa, 0x16, 0x2c, 0x4c, 0x12, 0x46, 0x20, 0x4, 0x56, 0x4e, 0x50, 0x5c, 0xe, 0x2e, 0x58,
			0x28, 0x2, 0x4a, 0x38, 0x30, 0x32, 0x3e, 0x54, 0x10, 0x3a, 0xa, 0x48, 0x2c, 0x1a, 0x12,
			0x14, 0x20, 0x36, 0x56, 0x1c, 0x50, 0x2a, 0xe, 0x60, 0x58, 0x5a, 0x2, 0x18, 0x38, 0x62,
			0x32, 0xc, 0x54, 0x42, 0x3a, 0x3c, 0x48, 0x5e, 0x1a, 0x44, 0x14, 0x52, 0x36, 0x24, 0x1c,
			0x1e, 0x2a, 0x40, 0x60, 0x26, 0x5a, 0x34, 0x18, 0x6, 0x62, 0x0, 0xc, 0x22, 0x42, 0x8,
			0x3c, 0x16, 0x5e, 0x4c, 0x44, 0x46, 0x52, 0x4, 0x24, 0x4e, 0x1e, 0x5c, 0x40, 0x2e, 0x26,
			0x28, 0x34, 0x4a, 0x6, 0x30, 0x0, 0x3e, 0x22, 0x10, 0x8, 0xa, 0x16, 0x2c, 0x4c, 0x12,
			0x46, 0x20, 0x4, 0x56, 0x4e, 0x50, 0x5c, 0xe, 0x2e, 0x58, 0x28, 0x2, 0x4a, 0x38, 0x30,
			0x32, 0x3e, 0x54, 0x10, 0x3a, 0xa, 0x48, 0x2c, 0x1a, 0x12, 0x14, 0x20, 0x36, 0x56, 0x1c,
			0x50, 0x2a, 0xe, 0x60, 0x58, 0x5a, 0x2, 0x18, 0x38, 0x62, 0x32, 0xc, 0x54, 0x42, 0x3a,
			0x3c, 0x48, 0x5e, 0x1a, 0x44, 0x14, 0x52, 0x36, 0x24, 0x1c, 0x1e, 0x2a, 0x40, 0x60, 0x26,
			0x5a, 0x34, 0x18, 0x6, 0x62, 0x0, 0xc, 0x22, 0x42, 0x8, 0x3c, 0x16, 0x5e, 0x4c, 0x44,
			0x46, 0x52, 0x4, 0x24, 0x4e, 0x1e, 0x5c, 0x40, 0x2e, 0x26, 0x28, 0x34, 0x4a, 0x6, 0x30,
			0x0, 0x3e, 0x22, 0x10, 0x8, 0xa, 0x16, 0x2c, 0x4c, 0x12, 0x46, 0x20, 0x4, 0x56, 0x4e,
			0x50, 0x5c, 0xe, 0x2e, 0x58, 0x28, 0x2, 0x4a, 0x38, 0x30, 0x32, 0x3e, 0x54, 0x10, 0x3a,
			0xa, 0x48, 0x2c, 0x1a, 0x12, 0x14, 0x20, 0x36, 0x56, 0x1c, 0x50, 0x2a, 0xe, 0x60, 0x58,
			0x5a, 0x2, 0x18, 0x38, 0x62, 0x32, 0xc, 0x54, 0x42, 0x3a, 0x3c, 0x48, 0x5e, 0x1a, 0x44,
			0x14, 0x52, 0x36, 0x24, 0x1c, 0x1e, 0x2a, 0x40, 0x60, 0x26, 0x5a, 0x34, 0x18, 0x6, 0x62,
			0x0, 0xc, 0x22, 0x42, 0x8, 0x3c, 0x16, 0x5e, 0x4c, 0x44, 0x46, 0x52, 0x4, 0x24, 0x4e,
			0x1e, 0x5c, 0x40, 0x2e, 0x26, 0x28, 0x34, 0x4a, 0x6, 0x30, 0x0, 0x3e, 0x22, 0x10, 0x8,
			0xa, 0x16, 0x2c, 0x4c, 0x12, 0x46, 0x20, 0x4, 0x56, 0x4e, 0x50, 0x5c, 0xe, 0x2e, 0x58,
			0x28, 0x2, 0x4a, 0x38, 0x30, 0x32, 0x3e, 0x54, 0x10, 0x3a, 0xa, 0x48, 0x2c, 0x1a, 0x12,
			0x14, 0x20, 0x36, 0x56, 0x1c, 0x50, 0x2a, 0xe, 0x60, 0x58, 0x5a, 0x2, 0x18, 0x38, 0x62,
			0x32, 0xc, 0x54, 0x42, 0x3a, 0x3c, 0x48, 0x5e, 0x1a, 0x44, 0x14, 0x52, 0x36, 0x24, 0x1c,
			0x1e, 0x2a, 0x40, 0x60, 0x26, 0x5a, 0x34, 0x18, 0x6, 0x62, 0x0, 0xc, 0x22, 0x42, 0x8,
			0x3c, 0x16, 0x5e, 0x4c, 0x44, 0x46, 0x52, 0x4, 0x24, 0x4e, 0x1e, 0x5c, 0x40, 0x2e, 0x26,
			0x28, 0x34, 0x4a, 0x6, 0x30, 0x0, 0x3e, 0x22, 0x10, 0x8, 0xa, 0x16, 0x2c, 0x4c, 0x12,
			0x46, 0x20, 0x4, 0x56, 0x4e, 0x50, 0x5c, 0xe, 0x2e, 0x58, 0x28, 0x2, 0x4a, 0x38, 0x30,
			0x32, 0x3e, 0x54, 0x10, 0x3a, 0xa, 0x48, 0x2c, 0x1a, 0x12, 0x14, 0x20, 0x36, 0x56, 0x1c,
			0x50, 0x2a, 0xe, 0x60, 0x58, 0x5a, 0x2, 0x18, 0x38, 0x62, 0x32, 0xc, 0x54, 0x42, 0x3a,
			0x3c, 0x48, 0x5e, 0x1a, 0x44, 0x14, 0x52, 0x36, 0x24, 0x1c, 0x1e, 0x2a, 0x40, 0x60, 0x26,
			0x5a, 0x34, 0x18, 0x6, 0x62, 0x0, 0xc, 0x22, 0x42, 0x8, 0x3c, 0x16, 0x5e, 0x4c, 0x44,
			0x46, 0x52, 0x4, 0x24, 0x4e, 0x1e, 0x5c, 0x40, 0x2e, 0x26, 0x28, 0x34, 0x4a, 0x6, 0x30,
			0x0, 0x3e, 0x22, 0x10, 0x8, 0xa, 0x16, 0x2c, 0x4c, 0x12, 0x46, 0x20, 0x4, 0x56, 0x4e,
			0x50, 0x5c, 0xe, 0x2e, 0x58, 0x28, 0x2, 0x4a, 0x38, 0x30, 0x32, 0x3e, 0x54, 0x10, 0x3a,
			0xa, 0x48, 0x2c, 0x1a, 0x12, 0x14, 0x20, 0x36, 0x56, 0x1c, 0x50, 0x2a, 0xe, 0x60, 0x58,
			0x5a, 0x2, 0x18, 0x38, 0x62, 0x32, 0xc, 0x54, 0x42, 0x3a, 0x3c, 0x48, 0x5e, 0x1a, 0x44,
			0x14, 0x52, 0x36, 0x24, 0x1c, 0x1e, 0x2a, 0x40, 0x60, 0x26, 0x5a, 0x34, 0x18, 0x6, 0x62,
			0x0, 0xc, 0x22, 0x42, 0x8, 0x3c, 0x16, 0x5e, 0x4c, 0x44, 0x46, 0x52, 0x4, 0x24, 0x4e,
			0x1e, 0x5c, 0x40, 0x2e, 0x26, 0x28, 0x34, 0x4a, 0x6, 0x30,
		},
		ShortTest: 32767,
	}
	if !reflect.DeepEqual(correct, l) {
		t.Log("Wanted: ", correct)
		t.Log("Got   : ", l)
		t.Fail()
	}
}