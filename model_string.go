// Code generated by "stringer -linecomment -type Model"; DO NOT EDIT.

package ptouchgo

import "strconv"

const (
	_Model_name_0 = "PT-P750W"
	_Model_name_1 = "PT-P710BT"
)

func (i Model) String() string {
	switch {
	case i == 104:
		return _Model_name_0
	case i == 118:
		return _Model_name_1
	default:
		return "Model(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
