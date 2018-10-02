// Code generated by go-enum
// DO NOT EDIT!

package data

import (
	"fmt"
)

const (
	// entity_typeType_0 is a entity_type of type Type_0
	entity_typeType_0 entity_type = iota
	// entity_typeType_1 is a entity_type of type Type_1
	entity_typeType_1
	// entity_typeType_2 is a entity_type of type Type_2
	entity_typeType_2
)

const _entity_typeName = "Type_0Type_1Type_2"

var _entity_typeMap = map[entity_type]string{
	0: _entity_typeName[0:6],
	1: _entity_typeName[6:12],
	2: _entity_typeName[12:18],
}

// String implements the Stringer interface.
func (x entity_type) String() string {
	if str, ok := _entity_typeMap[x]; ok {
		return str
	}
	return fmt.Sprintf("entity_type(%d)", x)
}

var _entity_typeValue = map[string]entity_type{
	_entity_typeName[0:6]:   0,
	_entity_typeName[6:12]:  1,
	_entity_typeName[12:18]: 2,
}

// Parseentity_type attempts to convert a string to a entity_type
func Parseentity_type(name string) (entity_type, error) {
	if x, ok := _entity_typeValue[name]; ok {
		return x, nil
	}
	return entity_type(0), fmt.Errorf("%s is not a valid entity_type", name)
}