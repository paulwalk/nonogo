package pkg

import (
	"encoding/binary"
	"strconv"
)

func getCellSlice(value, count int) []int {
	s := make([]int, 0)
	for i := 0; i < count; i++ {
		s = append(s, value)
	}
	return s
}

func sortLineLabel(name string) string {
	// split numeric suffix
	i := len(name) - 1
	for ; i >= 0; i-- {
		if '0' > name[i] || name[i] > '9' {
			break
		}
	}
	i++
	// string numeric suffix to uint64 bytes
	// empty string is zero, so integers are plus one
	b64 := make([]byte, 64/8)
	s64 := name[i:]
	if len(s64) > 0 {
		u64, err := strconv.ParseUint(s64, 10, 64)
		if err == nil {
			binary.BigEndian.PutUint64(b64, u64+1)
		}
	}
	return name[:i] + string(b64)
}
