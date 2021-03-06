package common

import (
	"time"
	"math/rand"
	"fmt"
)

const (
	letters = "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"
)

func init () {
	// TODO: save and replay seed
	rand.Seed(time.Now().Unix())
}

func RandInt(min int, max int) int {
	return rand.Intn(max - min) + min
}


func RandString(n int) string {
	var l int64 = int64(len(letters))
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = letters[int(rand.Int63n(l))]
	}
	return string(b)
}

func RandSlice(n int) []string {
	s := make([]string, n)
	for i:=range s {
		s[i] = RandString(16)
	}
	return s
}

func CheckFlag(flags uint32, flag int) bool {
	return 0 != (flags & uint32(flag))
}

func BytesToUint32(b []byte) uint32 {
	r := uint32(0)
	for i:=range b {
		r += uint32(b[i])
	}
	return r
}

func MakeGetAttrFlags(f ...int) []uint32 {
	r := []uint32{0}
	length := 1
	for i := range f {
		val := f[i]
		slot := val/32
		val = val % 32
		if slot > length-1 {
			for x:=0; x<=slot-length; x++ {
				r = append(r, 0)
			}
			length = len(r)
		}
		r[slot] |= (1<<uint32(val)) // ??
	}
	return r
}

func MakeUint32Flags(f ...int) uint32 {
	r := uint32(0)
	for i:=range f {
		r |= uint32(f[i])
	}
	return r
}

func GetBitmap(bits ...int) ([]uint32) {
	b := []uint32{0,0}
	// it will panic in case of bit > 64
	for _, v := range bits {
		b[v/32] |= (1 << uint32(v%32))
	}
	return b
}

func GetPermAttrList(perm uint) (l []byte) {
	l = make([]byte, 4)
	l[3] = byte(perm & 0xff)
	l[2] = byte((perm & 0xff00) >> 8)
	l[1] = byte((perm & 0xff0000) >> 16)
	l[0] = byte((perm & 0xff000000) >> 24)
	return l
}

// TODO : polymorphic
func InSliceUint32(val uint32, slice []uint32) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

func Tm() string {
	t := time.Now()
	return fmt.Sprintf("[%02d:%02d:%02d]",
		t.Hour(), t.Minute(), t.Second())
}

