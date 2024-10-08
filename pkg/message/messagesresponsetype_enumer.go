// Code generated by "enumer -type=MessagesResponseType -linecomment -json=true -text=true"; DO NOT EDIT.

package message

import (
	"encoding/json"
	"fmt"
	"strings"
)

const _MessagesResponseTypeName = "messageerror"

var _MessagesResponseTypeIndex = [...]uint8{0, 7, 12}

const _MessagesResponseTypeLowerName = "messageerror"

func (i MessagesResponseType) String() string {
	if i < 0 || i >= MessagesResponseType(len(_MessagesResponseTypeIndex)-1) {
		return fmt.Sprintf("MessagesResponseType(%d)", i)
	}
	return _MessagesResponseTypeName[_MessagesResponseTypeIndex[i]:_MessagesResponseTypeIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _MessagesResponseTypeNoOp() {
	var x [1]struct{}
	_ = x[MessagesResponseTypeMessage-(0)]
	_ = x[MessagesResponseTypeError-(1)]
}

var _MessagesResponseTypeValues = []MessagesResponseType{MessagesResponseTypeMessage, MessagesResponseTypeError}

var _MessagesResponseTypeNameToValueMap = map[string]MessagesResponseType{
	_MessagesResponseTypeName[0:7]:       MessagesResponseTypeMessage,
	_MessagesResponseTypeLowerName[0:7]:  MessagesResponseTypeMessage,
	_MessagesResponseTypeName[7:12]:      MessagesResponseTypeError,
	_MessagesResponseTypeLowerName[7:12]: MessagesResponseTypeError,
}

var _MessagesResponseTypeNames = []string{
	_MessagesResponseTypeName[0:7],
	_MessagesResponseTypeName[7:12],
}

// MessagesResponseTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func MessagesResponseTypeString(s string) (MessagesResponseType, error) {
	if val, ok := _MessagesResponseTypeNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _MessagesResponseTypeNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to MessagesResponseType values", s)
}

// MessagesResponseTypeValues returns all values of the enum
func MessagesResponseTypeValues() []MessagesResponseType {
	return _MessagesResponseTypeValues
}

// MessagesResponseTypeStrings returns a slice of all String values of the enum
func MessagesResponseTypeStrings() []string {
	strs := make([]string, len(_MessagesResponseTypeNames))
	copy(strs, _MessagesResponseTypeNames)
	return strs
}

// IsAMessagesResponseType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i MessagesResponseType) IsAMessagesResponseType() bool {
	for _, v := range _MessagesResponseTypeValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for MessagesResponseType
func (i MessagesResponseType) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for MessagesResponseType
func (i *MessagesResponseType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("MessagesResponseType should be a string, got %s", data)
	}

	var err error
	*i, err = MessagesResponseTypeString(s)
	return err
}

// MarshalText implements the encoding.TextMarshaler interface for MessagesResponseType
func (i MessagesResponseType) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for MessagesResponseType
func (i *MessagesResponseType) UnmarshalText(text []byte) error {
	var err error
	*i, err = MessagesResponseTypeString(string(text))
	return err
}
