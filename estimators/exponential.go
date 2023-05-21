package estimators

import (
	"math/bits"
	"sync/atomic"
)

type ExponentialEstimator struct {
	Buckets                       int
	AllowedMemoryGrowthPercentage int
}

var _ Estimator = ExponentialEstimator{}

func (e ExponentialEstimator) Size() int { return e.Buckets + 1 }
func (e ExponentialEstimator) Init(Hint) {}
func (e ExponentialEstimator) Add(hint Hint, size uint32) {
	atomic.AddUint64(&hint[0], 1)
	bucket := 1 + (32 - bits.LeadingZeros32(size))
	if bucket >= len(hint) {
		bucket = len(hint) - 1
	}
	atomic.AddUint64(&hint[bucket], 1)
}
func (e ExponentialEstimator) Estimate(hint Hint) uint32 {
	current, total := uint64(0), atomic.LoadUint64(&hint[0])
	for i := 1; i < len(hint); i++ {
		current += atomic.LoadUint64(&hint[i])
		if 100*current > uint64(e.AllowedMemoryGrowthPercentage)*total {
			return (1 << (i - 1)) - 1
		}
	}
	return 0
}
