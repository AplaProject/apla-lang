package runtime

import (
	"fmt"
	"strings"
)

// Replace replaces old substrings to new substrings
func Replace(rt *Runtime, s, old, new int64) int64 {
	rt.Strings = append(rt.Strings, strings.Replace(rt.Strings[s], rt.Strings[old],
		rt.Strings[new], -1))
	return int64(len(rt.Strings) - 1)
}

// Split splits the input string to array
func Split(rt *Runtime, input, sep int64) int64 {
	out := strings.Split(rt.Strings[input], rt.Strings[sep])
	result := make([]int64, len(out))
	for i, val := range out {
		rt.Strings = append(rt.Strings, val)
		result[i] = int64(len(rt.Strings) - 1)
	}
	rt.Objects = append(rt.Objects, result)
	return int64(len(rt.Objects) - 1)
}

// Substr returns a substring with the specified offset and length
func Substr(rt *Runtime, in, off, length int64) (int64, error) {
	var rin []rune
	rin = []rune(rt.Strings[in])
	rlen := int64(len(rin))
	if length < 0 {
		length = -length
		off -= length
	}
	if off < 0 || off >= rlen || off+length > rlen {
		return 0, fmt.Errorf(errInvalidParam)
	}
	if length == 0 {
		length = rlen - off
	}
	rt.Strings = append(rt.Strings, string(rin[off:off+length]))
	return int64(len(rt.Strings) - 1), nil
}
