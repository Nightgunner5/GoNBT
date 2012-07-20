package main

import (
	"os"
	"fmt"
	"errhandler"
	"encoding/json"
	"github.com/bemasher/GoNBT"
)

const (
	NBTFILE = "bigtest.nbt"
)

type Level struct {
	Nested Nested `nbt:"nested compound test"`
	ByteTest byte
	IntTest int32
	StringTest string
	ListTestLong []int64 `nbt:"listTest (long)"`
	DoubleTest float64
	FloatTest float32
	LongTest int64
	ListTestCompound []ListTest `nbt:"listtest (compound)"`
	ByteArrayTest []byte `nbt:"bytearraytest (the first 1000 values of (n*n*255+n*7)%100, starting with n=0 (0, 62, 34, 16, 8, ...))"`
	ShortTest int16
}

type Nested struct {
	Egg Egg
	Ham Ham
}

type Egg struct {
	Name string
	Value float32
}

type Ham struct {
	Name string
	Value float32
}

type ListTest struct {
	CreatedOn int64 `nbt:"created-on"`
	Name string
}

func main() {
	nbtFile, err := os.Open(NBTFILE)
	errhandler.Handle("Error opening chunk file: ", err)
	defer nbtFile.Close()
	
	var l Level
	nbt.Read(nbtFile, &l)
	
	data, err := json.MarshalIndent(l, "", "\t")
	errhandler.Handle("Error marshalling json: ", err)
	
	fmt.Printf("%s\n", data)
}