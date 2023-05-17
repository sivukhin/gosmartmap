package smartmap

import (
	"math/bits"
	"math/rand"
	"runtime"
	"sync"
)

const (
	SizeHintBuckets        = 16
	MemoryPercentThreshold = 90
)

type SizeHint struct {
	Pc      uintptr
	Total   uint64
	Buckets [SizeHintBuckets]uint64
}

func (h *SizeHint) estimate() int {
	small := uint64(0)
	for i := 0; i < SizeHintBuckets; i++ {
		small += h.Buckets[i]
		if 100*small > MemoryPercentThreshold*h.Total {
			return (1 << i) - 1
		}
	}
	return 0
}

type container[TKey comparable, TValue any] struct {
	m    map[TKey]TValue
	hint *SizeHint
}

var (
	hintsLock sync.Mutex

	hintsPower        = 10
	hintsSizeMinusOne = uint32(1<<hintsPower) - 1
	hints             = make([]*SizeHint, hintsSizeMinusOne+1)
	hintsHash         = rand.Uint32() | 1
	hintsCount        = uint32(0)
)

func slot(hs []*SizeHint, pc uintptr) (uint32, bool) {
	h := (hintsHash * uint32(pc)) >> (32 - hintsPower)
	for {
		if hs[h] == nil {
			return h, false
		}
		if hs[h].Pc == pc {
			return h, true
		}
		h = (h + 1) & (hintsSizeMinusOne)
	}
}

func grow() {
	next := make([]*SizeHint, 1<<(hintsPower+1))
	size := 1 << hintsPower
	for i := 0; i < size; i++ {
		if hints[i] != nil {
			position, _ := slot(next, hints[i].Pc)
			next[position] = hints[i]
		}
	}

	hintsPower++
	hintsSizeMinusOne = (1 << hintsPower) - 1
	hints = next
}

func Make[TKey comparable, TValue any]() map[TKey]TValue {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		return make(map[TKey]TValue)
	}

	if hintsCount > hintsSizeMinusOne/2 {
		hintsLock.Lock()
		grow()
		hintsLock.Unlock()
	}

	position, ok := slot(hints, pc)
	if !ok {
		hintsCount++
		hints[position] = new(SizeHint)
		hints[position].Pc = pc
	}

	return MakeHint[TKey, TValue](hints[position])
}

func MakeHint[TKey comparable, TValue any](hint *SizeHint) map[TKey]TValue {
	size := hint.estimate()
	m := make(map[TKey]TValue, size)
	c := &container[TKey, TValue]{m: m, hint: hint}
	runtime.SetFinalizer(c, func(c *container[TKey, TValue]) {
		size := uint32(len(c.m))
		bucket := 32 - bits.LeadingZeros32(size)
		if bucket >= SizeHintBuckets {
			bucket = SizeHintBuckets - 1
		}
		c.hint.Total++
		c.hint.Buckets[bucket]++
	})
	return m
}
