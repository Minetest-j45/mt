// Code generated by "stringer -trimprefix Style -type HUDStyleFlags"; DO NOT EDIT.

package mt

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[StyleBold-1]
	_ = x[StyleItalic-2]
	_ = x[StyleMono-4]
}

const (
	_HUDStyleFlags_name_0 = "BoldItalic"
	_HUDStyleFlags_name_1 = "Mono"
)

var (
	_HUDStyleFlags_index_0 = [...]uint8{0, 4, 10}
)

func (i HUDStyleFlags) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _HUDStyleFlags_name_0[_HUDStyleFlags_index_0[i]:_HUDStyleFlags_index_0[i+1]]
	case i == 4:
		return _HUDStyleFlags_name_1
	default:
		return "HUDStyleFlags(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
