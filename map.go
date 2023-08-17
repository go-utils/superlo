package superlo

import "github.com/samber/lo"

// Map - returns a new collection of mapped values and error
func Map[T any, R any](collection []T, iteratee func(item T, index int) (R, error)) (mappedCollection []R, topErr error) {
	defer func() {
		_ = recover()
	}()
	return lo.Map(collection, func(item T, index int) R {
		returnValue, err := iteratee(item, index)
		if err != nil {
			topErr = err
			panic(nil)
		}
		return returnValue
	}), nil
}
