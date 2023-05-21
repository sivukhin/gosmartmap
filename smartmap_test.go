package smartmap

import (
	"github.com/stretchr/testify/require"
	"math/rand"
	"strconv"
	"testing"
)

func BenchmarkMap(b *testing.B) {
	const maxSize = 1024 * 1024
	keys := make([]string, maxSize)
	values := make([]string, maxSize)
	for i := 0; i < maxSize; i++ {
		keys[i] = strconv.Itoa(rand.Int())
		values[i] = strconv.Itoa(rand.Int())
	}
	b.Run("smartmap-0", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			m := Make[string, string]()
			require.Len(b, m, 0)
		}
	})
	b.Run("smartmap-128", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			size := rand.Intn(128 + 1)
			m := Make[string, string]()
			for s := 0; s < size; s++ {
				m[keys[s]] = values[s]
			}
			require.Len(b, m, size)
		}
	})
	b.Run("smartmap-1024", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			size := rand.Intn(1024 + 1)
			m := Make[string, string]()
			for s := 0; s < size; s++ {
				m[keys[s]] = values[s]
			}
			require.Len(b, m, size)
		}
	})
	b.Run("smartmap-65536", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			size := rand.Intn(65536 + 1)
			m := Make[string, string]()
			for s := 0; s < size; s++ {
				m[keys[s]] = values[s]
			}
			require.Len(b, m, size)
		}
	})
}
