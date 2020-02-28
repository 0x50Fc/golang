package top

import "time"

func NewSID(rate int32) int64 {
	tv := time.Now().UnixNano() / int64(time.Millisecond)
	return (int64(rate) << 32) | (tv & 0x0ffffffff)
}

func NewSIDWithTimestamp(rate int32, timestamp int64) int64 {
	return (int64(rate) << 32) | (timestamp & 0x0ffffffff)
}
