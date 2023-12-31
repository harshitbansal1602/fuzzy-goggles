// Code generated by "enumer -type=Lifestyle -json"; DO NOT EDIT.

package handlers

import (
	"encoding/json"
	"fmt"
)

const _LifestyleName = "LongDrivesFoodieImpulsiveBuyerBingeWatcherLoveGadgets"

var _LifestyleIndex = [...]uint8{0, 10, 16, 30, 42, 53}

func (i Lifestyle) String() string {
	i -= 1
	if i < 0 || i >= Lifestyle(len(_LifestyleIndex)-1) {
		return fmt.Sprintf("Lifestyle(%d)", i+1)
	}
	return _LifestyleName[_LifestyleIndex[i]:_LifestyleIndex[i+1]]
}

var _LifestyleValues = []Lifestyle{1, 2, 3, 4, 5}

var _LifestyleNameToValueMap = map[string]Lifestyle{
	_LifestyleName[0:10]:  1,
	_LifestyleName[10:16]: 2,
	_LifestyleName[16:30]: 3,
	_LifestyleName[30:42]: 4,
	_LifestyleName[42:53]: 5,
}

// LifestyleString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func LifestyleString(s string) (Lifestyle, error) {
	if val, ok := _LifestyleNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Lifestyle values", s)
}

// LifestyleValues returns all values of the enum
func LifestyleValues() []Lifestyle {
	return _LifestyleValues
}

// IsALifestyle returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Lifestyle) IsALifestyle() bool {
	for _, v := range _LifestyleValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for Lifestyle
func (i Lifestyle) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for Lifestyle
func (i *Lifestyle) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Lifestyle should be a string, got %s", data)
	}

	var err error
	*i, err = LifestyleString(s)
	return err
}
