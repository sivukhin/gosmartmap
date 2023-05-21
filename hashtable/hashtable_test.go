//go:build !race

package hashtable

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestSinglePut(t *testing.T) {
	h := NewHashTable(0)
	_, ok := h.Get(1)
	require.Equal(t, 0, h.Size())
	require.False(t, ok)
	require.Len(t, h.Set(1, 1), 1)
	require.Equal(t, 1, h.Size())
}

func TestManyPuts(t *testing.T) {
	h := NewHashTable(0)
	size := 1 << 14
	for i := 0; i < size; i++ {
		buffer := h.Set(uint64(i), 1)
		buffer[0] = uint64(i)
	}
	require.Equal(t, size, h.Size())
	for i := 0; i < size; i++ {
		buffer, ok := h.Get(uint64(i))
		require.True(t, ok)
		require.Equal(t, uint64(i), buffer[0])
	}
}

func TestConcurrentAccess(t *testing.T) {
	h := NewHashTable(0)
	wg := sync.WaitGroup{}
	for i := 0; i < 1<<12; i++ {
		h.Set(uint64(i), 1)
	}
	start := time.Now()
	for i := 0; i < 128; i++ {
		wg.Add(1)
		go func() {
			for s := 0; s < 1<<14; s++ {
				key := rand.Uint64() % (1 << 12)
				buffer, ok := h.Get(key)
				if !ok || rand.Intn(2) == 0 {
					_ = h.Set(key, 1)
				} else {
					require.Len(t, buffer, 1)
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("elapsed: %v\n", time.Since(start))
}
