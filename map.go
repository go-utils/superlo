package superlo

import (
	"sync"

	"golang.org/x/sync/errgroup"
)

// Map - returns a new collection of mapped values and error
func Map[T any, R any](collection []T, iteratee func(item T, index int) (R, error)) ([]R, error) {
	result := make([]R, len(collection))

	for i, item := range collection {
		returnValue, err := iteratee(item, i)
		if err != nil {
			return nil, err
		}
		result[i] = returnValue
	}

	return result, nil
}

// ParallelMap manipulates a slice and transforms it to a slice of another type.
// `iteratee` is call in parallel. Result keep the same order.
func ParallelMap[T any, R any](collection []T, iteratee func(item T, index int) (R, error)) ([]R, error) {
	result := make([]R, len(collection))

	var (
		eg errgroup.Group
		mu sync.Mutex
	)

	for i, item := range collection {
		func(_item T, _i int) {
			eg.Go(func() error {
				returnValue, err := iteratee(_item, _i)
				if err != nil {
					return err
				}

				mu.Lock()
				result[_i] = returnValue
				mu.Unlock()

				return nil
			})
		}(item, i)
	}

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	return result, nil
}
