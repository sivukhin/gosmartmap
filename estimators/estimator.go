package estimators

type (
	Hint      []uint64
	Estimator interface {
		Size() int
		Init(hint Hint)
		Add(hint Hint, size uint32)
		Estimate(hint Hint) uint32
	}
)
