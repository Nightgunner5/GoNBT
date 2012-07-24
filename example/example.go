package main

import (
	"encoding/json"
	"fmt"
	"github.com/bemasher/errhandler"
	"nbt"
	"os"
)

const (
	NBTFILE = "bigtest.nbt"
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

func main() {
	nbtFile, err := os.Open(NBTFILE)
	errhandler.Handle("Error opening chunk file: ", err)
	defer nbtFile.Close()

	var l Level
	nbt.Read(nbtFile, &l)

	// Using json for pretty indented printing of the structure
	data, err := json.MarshalIndent(l, "", "\t")
	errhandler.Handle("Error marshalling json: ", err)
	fmt.Printf("%s\n", data)
}
