// Code generated by "stringer -trimprefix P2 -type Param2Type"; DO NOT EDIT.

package mt

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[P2Nibble-0]
	_ = x[P2Byte-1]
	_ = x[P2Flowing-2]
	_ = x[P2FaceDir-3]
	_ = x[P2Mounted-4]
	_ = x[P2Leveled-5]
	_ = x[P2Rotation-6]
	_ = x[P2Mesh-7]
	_ = x[P2Color-8]
	_ = x[P2ColorFaceDir-9]
	_ = x[P2ColorMounted-10]
	_ = x[P2GlassLikeLevel-11]
	_ = x[P2ColorRotation-12]
}

const _Param2Type_name = "NibbleByteFlowingFaceDirMountedLeveledRotationMeshColorColorFaceDirColorMountedGlassLikeLevelColorRotation"

var _Param2Type_index = [...]uint8{0, 6, 10, 17, 24, 31, 38, 46, 50, 55, 67, 79, 93, 106}

func (i Param2Type) String() string {
	if i >= Param2Type(len(_Param2Type_index)-1) {
		return "Param2Type(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Param2Type_name[_Param2Type_index[i]:_Param2Type_index[i+1]]
}
