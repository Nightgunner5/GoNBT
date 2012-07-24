package nbt

import (
	"encoding/binary"
	"github.com/bemasher/errhandler"
	"io"
	"reflect"
	"strings"
)

// Define constants defining all of the possible tag types
const (
	TAG_END = iota
	TAG_BYTE
	TAG_SHORT
	TAG_INT
	TAG_LONG
	TAG_FLOAT
	TAG_DOUBLE
	TAG_BYTE_ARRAY
	TAG_STRING
	TAG_LIST
	TAG_COMPOUND
	TAG_INT_ARRAY
	TAG_UNKNOWN
)

// This type represents an NBT tag's type.
type TagType byte

// Provides a string representation of the tag type.
func (t TagType) String() string {
	switch t {
	case TAG_END:
		return "TAG_END"
	case TAG_BYTE:
		return "TAG_BYTE"
	case TAG_SHORT:
		return "TAG_SHORT"
	case TAG_INT:
		return "TAG_INT"
	case TAG_LONG:
		return "TAG_LONG"
	case TAG_FLOAT:
		return "TAG_FLOAT"
	case TAG_DOUBLE:
		return "TAG_DOUBLE"
	case TAG_BYTE_ARRAY:
		return "TAG_BYTE_ARRAY"
	case TAG_STRING:
		return "TAG_STRING"
	case TAG_LIST:
		return "TAG_LIST"
	case TAG_COMPOUND:
		return "TAG_COMPOUND"
	case TAG_INT_ARRAY:
		return "TAG_INT_ARRAY"
	}
	return "TAG_UNKNOWN"
}

// Using reflection create a map of the names of fields (or the nbt tag if present) to reflect.Value for a given struct.
func StructFields(d reflect.Value) map[string]reflect.Value {
	// Make sure this is a struct
	if d.Kind() == reflect.Struct {
		tags := make(map[string]reflect.Value, 0)

		// Include the type's name as well, this is used when decoding compound tags
		tags[strings.ToLower(d.Type().Name())] = d

		// For each of the fields in the struct
		for i := 0; i < d.NumField(); i++ {
			// If the nbt key is present in the field's tags use that instead otherwise just use the field name
			if field := d.Type().Field(i); field.Tag.Get("nbt") == "" {
				tags[strings.ToLower(field.Name)] = d.Field(i)
			} else {
				tags[strings.ToLower(field.Tag.Get("nbt"))] = d.Field(i)
			}
		}
		return tags
	}
	return nil
}

// Reads NBT formatted data into the given interface. Data must be a pointer to
// a variable of kind struct. Field order in the struct is not important, fields
// will be written when an NBT is found with a matching name. Embedded structs
// will be descended into, this corresponds to compound tags.
// The caller is responsible for decompressing the data before calling this.
func Read(r io.Reader, data interface{}) {
	readTag(r, TAG_UNKNOWN, reflect.ValueOf(data).Elem())
}

func readTag(r io.Reader, tagType TagType, data reflect.Value) (name string, newTagType TagType) {
	// If TAG_UNKNOWN is passed we're reading a full tag which is not a member of a list
	if tagType == TAG_UNKNOWN {
		newTagType = TagType(ReadByte(r))
		// If we arrived here after being called by readCompound we should have a TAG_END so just return
		if newTagType == TAG_END {
			return
		}
		// If we get here then this tag should have a name
		name = ReadString(r)
	} else {
		newTagType = TagType(tagType)
	}

	// Get a list of the fields for the given struct
	var field reflect.Value
	fields := StructFields(data)

	// If this field is valid it isn't a zero value so assign it's field from the struct map
	if data.IsValid() {
		field = fields[strings.ToLower(name)]
	}

	// If this is the root tag keep going anyway and just decode into the current value
	if name == "" {
		field = data
	}

	var (
		temp interface{}
		list reflect.Value
	)
	switch newTagType {
	case TAG_BYTE:
		temp = ReadByte(r)
	case TAG_SHORT:
		temp = ReadShort(r)
	case TAG_INT:
		temp = ReadInt(r)
	case TAG_LONG:
		temp = ReadLong(r)
	case TAG_FLOAT:
		temp = ReadFloat(r)
	case TAG_DOUBLE:
		temp = ReadDouble(r)
	case TAG_BYTE_ARRAY:
		temp = ReadByteArray(r)
	case TAG_STRING:
		temp = ReadString(r)
	case TAG_INT_ARRAY:
		temp = ReadIntArray(r)
	case TAG_COMPOUND:
		ReadCompound(r, field)
	case TAG_LIST:
		list = ReadList(r, field)
	}

	// If field was invalid we still needed to decode the tag, if we got here and it's valid we need to store it
	if field.IsValid() {
		// If temp has a value in it, then it wasn't a list or compound tag so go ahead and store it
		if temp != nil {
			v := reflect.ValueOf(temp)
			if field.Type().AssignableTo(v.Type()) {
				field.Set(v)
			}
			// If the list valud is valid, make sure we store it
		} else if list.IsValid() {
			field.Set(list)
		}
	}
	return
}

// Reads a single byte.
func ReadByte(r io.Reader) (i byte) {
	err := binary.Read(r, binary.BigEndian, &i)
	errhandler.Handle("Error reading Byte: ", err)
	return
}

// Reads a signed 16-bit integer.
func ReadShort(r io.Reader) (i int16) {
	err := binary.Read(r, binary.BigEndian, &i)
	errhandler.Handle("Error reading Short: ", err)
	return
}

// Reads a signed 32-bit integer.
func ReadInt(r io.Reader) (i int32) {
	err := binary.Read(r, binary.BigEndian, &i)
	errhandler.Handle("Error reading Int: ", err)
	return
}

// Reads a signed 64-bit integer.
func ReadLong(r io.Reader) (i int64) {
	err := binary.Read(r, binary.BigEndian, &i)
	errhandler.Handle("Error reading Long: ", err)
	return
}

// Reads a 32-bit IEEE-754 floating point value.
func ReadFloat(r io.Reader) (i float32) {
	err := binary.Read(r, binary.BigEndian, &i)
	errhandler.Handle("Error reading Float: ", err)
	return
}

// Reads a 64-bit IEEE-754 floating point value.
func ReadDouble(r io.Reader) (i float64) {
	err := binary.Read(r, binary.BigEndian, &i)
	errhandler.Handle("Error reading Double: ", err)
	return
}

// Reads a byte array, the array length (signed 32-bit int) is read first, then the array of bytes.
func ReadByteArray(r io.Reader) (i []byte) {
	i = make([]byte, ReadInt(r))
	_, err := r.Read(i)
	errhandler.Handle("Error reading Byte Array: ", err)
	return
}

// Reads a string, the length of the string is read first (signed 16-bit int), then the string characters.
func ReadString(r io.Reader) string {
	result := make([]byte, ReadShort(r))
	_, err := r.Read(result)
	errhandler.Handle("Error reading String: ", err)
	return string(result)
}

// Reads an array of 32-bit ints, the array length (signed 32-bit int) is read first, then the array of ints.
func ReadIntArray(r io.Reader) (list []int32) {
	length := int(ReadInt(r))
	list = make([]int32, 0, length)
	for i := 0; i < length; i++ {
		list = append(list, ReadInt(r))
	}
	return
}

// Reads a list of tags into the given data structure until a tag of TAG_END is encountered.
func ReadCompound(r io.Reader, data reflect.Value) {
	tagType := TAG_UNKNOWN
	for tagType != TAG_END {
		_, tagType = readTag(r, TAG_UNKNOWN, data)
	}
}

// Reads a list of tags into a list, the elements of the list don't have names or types specified, so readTag requires the tag type to be given.
func ReadList(r io.Reader, data reflect.Value) reflect.Value {
	listType := TagType(ReadByte(r))
	length := ReadInt(r)
	for i := int32(0); i < length; i++ {
		if data.IsValid() {
			temp := reflect.Indirect(reflect.New(data.Type().Elem()))
			readTag(r, listType, temp)
			data = reflect.Append(data, temp)
		} else {
			readTag(r, listType, data)
		}
	}
	return data
}
