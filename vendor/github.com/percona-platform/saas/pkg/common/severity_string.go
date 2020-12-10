// Code generated by "stringer -type=Severity -linecomment"; DO NOT EDIT.

package common

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Unknown-0]
	_ = x[Emergency-1]
	_ = x[Alert-2]
	_ = x[Critical-3]
	_ = x[Error-4]
	_ = x[Warning-5]
	_ = x[Notice-6]
	_ = x[Info-7]
	_ = x[Debug-8]
}

const _Severity_name = "unknownemergencyalertcriticalerrorwarningnoticeinfodebug"

var _Severity_index = [...]uint8{0, 7, 16, 21, 29, 34, 41, 47, 51, 56}

func (i Severity) String() string {
	if i < 0 || i >= Severity(len(_Severity_index)-1) {
		return "Severity(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Severity_name[_Severity_index[i]:_Severity_index[i+1]]
}
