package smartmap

import (
	"github.com/sivukhin/gosmartmap/estimators"
	"github.com/sivukhin/gosmartmap/hashtable"
	"runtime"
)

type Context struct {
	Estimator estimators.Estimator
	HashTable hashtable.HashTable
}

var DefaultContext = Context{
	Estimator: estimators.ExponentialEstimator{
		AllowedMemoryGrowthPercentage: 10,
		Buckets:                       16,
	},
	HashTable: hashtable.NewHashTable(10),
}

type container[TKey comparable, TValue any] struct {
	m         map[TKey]TValue
	estimator estimators.Estimator
	hint      []uint64
}

func (c Context) NewHint() estimators.Hint {
	hint := make([]uint64, c.Estimator.Size())
	c.Estimator.Init(hint)
	return hint
}

// Make creates a map with size estimated from statistics based on previous calls from same caller
// Make use information from runtime.Callers in order to identify different call-sites
// Make accept "optional" Context parameter with specified values for size estimator and hashtable
// go:noinline
func Make[TKey comparable, TValue any](optionalContext ...Context) map[TKey]TValue {
	context := DefaultContext
	if len(optionalContext) > 0 {
		context = optionalContext[0]
	}
	var pcs [1]uintptr
	n := runtime.Callers(2, pcs[:])
	if n == 0 {
		return make(map[TKey]TValue, 0)
	}
	pc := pcs[0]

	hint, found := context.HashTable.Get(uint64(pc))
	if !found {
		hint = context.HashTable.Set(uint64(pc), context.Estimator.Size())
		context.Estimator.Init(hint)
	}
	return MakeHint[TKey, TValue](hint, context)
}

// MakeHint creates a map with size estimated from a hint with accumulated statistics based on previous calls from same caller
// MakeHint accepts hint container which should be created with Context.NewHint method
func MakeHint[TKey comparable, TValue any](hint estimators.Hint, optionalContext ...Context) map[TKey]TValue {
	context := DefaultContext
	if len(optionalContext) > 0 {
		context = optionalContext[0]
	}
	size := context.Estimator.Estimate(hint)
	m := make(map[TKey]TValue, size)
	c := &container[TKey, TValue]{
		m:         m,
		estimator: context.Estimator,
		hint:      hint,
	}
	runtime.SetFinalizer(c, func(c *container[TKey, TValue]) { c.estimator.Add(c.hint, uint32(len(c.m))) })
	return m
}
