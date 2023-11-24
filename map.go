package superlo

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
