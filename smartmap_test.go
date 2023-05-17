package smartmap

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"math/rand"
	"strconv"
	"testing"
)

func BenchmarkMap(b *testing.B) {
	for _, size := range []int{128, 1024, 65536} {
		keys := make([]string, size)
		values := make([]string, size)
		for i := 0; i < len(keys); i++ {
			keys[i] = strconv.Itoa(rand.Int())
			values[i] = strconv.Itoa(rand.Int())
		}

		b.Run(fmt.Sprintf("map-%v", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				size := rand.Intn(size)
				m := Make[string, string]()
				for s := 0; s < size; s++ {
					m[keys[s]] = values[s]
				}
				require.Len(b, m, size)
			}
		})
	}
}
