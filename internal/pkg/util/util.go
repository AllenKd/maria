package util

import (
	"context"
	"sort"
	"sync"
	"time"
)

func MapKeys(m map[string]func(group *sync.WaitGroup, ctx context.Context, delay time.Duration)) (keys []string) {

	keys = make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}

	sort.Strings(keys)
	return
}
