package lookin

import (
	"github.com/hailongz/golang/basex"
)

func ByteCount(v int64) int {
	n := 0
	for v != 0 && n < 8 {
		v = v >> 8
		n = n + 1
	}
	if n == 0 {
		return 1
	}
	return n
}

func Encode(ids []int64) []byte {
	byteCount := 0
	for _, id := range ids {
		n := ByteCount(id)
		if n > byteCount {
			byteCount = n
		}
	}
	dst := make([]byte, 1+len(ids)*byteCount)
	i := 0
	dst[i] = (byte)(byteCount & 0x07f)
	i += 1
	for _, id := range ids {
		n := byteCount
		for n > 0 {
			dst[i] = (byte)(id & 0x0ff)
			n -= 1
			id = id >> 8
			i += 1
		}
	}
	return dst
}

func EncodeToString(ids []int64) string {
	return basex.Base62.Encode(Encode(ids))
}

func Decode(b []byte) []int64 {
	if b == nil {
		return nil
	}
	n := len(b)
	if n <= 1 {
		return []int64{}
	}
	byteCount := (int)(b[0])
	if byteCount > 8 || byteCount < 1 {
		return []int64{}
	}
	c := (n - 1) / byteCount
	dst := make([]int64, c)
	var v int64 = 0
	for i := 0; i < c; i++ {
		v = 0
		for j := i*byteCount + byteCount; j > i*byteCount; j-- {
			v = (v << 8) | (int64(b[j]) & 0x0ff)
		}
		dst[i] = v
	}
	return dst
}

func DeocdeString(s string) ([]int64, error) {
	b, err := basex.Base62.Decode(s)
	if err != nil {
		return nil, err
	}
	return Decode(b), nil
}
